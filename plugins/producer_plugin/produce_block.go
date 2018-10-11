package producer_plugin

import (
	"fmt"
	"github.com/eosspark/eos-go/chain/types"
	"github.com/eosspark/eos-go/common"
	"github.com/eosspark/eos-go/crypto"
	"github.com/eosspark/eos-go/crypto/ecc"
	"github.com/eosspark/eos-go/log"
	Chain "github.com/eosspark/eos-go/plugins/producer_plugin/mock"
	. "github.com/eosspark/eos-go/exception"
	"github.com/eosspark/eos-go/exception/try"
)



func (impl *ProducerPluginImpl) CalculateNextBlockTime(producerName common.AccountName, currentBlockTime common.BlockTimeStamp) *common.TimePoint {
	var result common.TimePoint

	chain := Chain.GetControllerInstance()

	hbs := chain.HeadBlockState()
	activeSchedule := hbs.ActiveSchedule.Producers

	// determine if this producer is in the active schedule and if so, where
	var itr *types.ProducerKey
	var producerIndex uint32
	for index, asp := range activeSchedule {
		if asp.AccountName == producerName {
			itr = &asp
			producerIndex = uint32(index)
			break
		}
	}

	if itr == nil {
		// this producer is not in the active producer set
		return nil
	}

	var minOffset uint32 = 1 // must at least be the "next" block

	// account for a watermark in the future which is disqualifying this producer for now
	// this is conservative assuming no blocks are dropped.  If blocks are dropped the watermark will
	// disqualify this producer for longer but it is assumed they will wake up, determine that they
	// are disqualified for longer due to skipped blocks and re-caculate their next block with better
	// information then
	currentWatermark, hasCurrentWatermark := impl.ProducerWatermarks[producerName]
	if hasCurrentWatermark {
		blockNum := chain.PendingBlockState().BlockNum
		if chain.PendingBlockState() != nil {
			blockNum++
		}
		if currentWatermark > blockNum {
			minOffset = currentWatermark - blockNum + 1
		}
	}

	// this producers next opportuity to produce is the next time its slot arrives after or at the calculated minimum
	minSlot := uint32(currentBlockTime) + minOffset
	minSlotProducerIndex := (minSlot % (uint32(len(activeSchedule)) * uint32(common.DefaultConfig.ProducerRepetitions))) / uint32(common.DefaultConfig.ProducerRepetitions)
	if producerIndex == minSlotProducerIndex {
		// this is the producer for the minimum slot, go with that
		result = common.BlockTimeStamp(minSlot).ToTimePoint()
	} else {
		// calculate how many rounds are between the minimum producer and the producer in question
		producerDistance := producerIndex - minSlotProducerIndex
		// check for unsigned underflow
		if producerDistance > producerIndex {
			producerDistance += uint32(len(activeSchedule))
		}

		// align the minimum slot to the first of its set of reps
		firstMinProducerSlot := minSlot - (minSlot % uint32(common.DefaultConfig.ProducerRepetitions))

		// offset the aligned minimum to the *earliest* next set of slots for this producer
		nextBlockSlot := firstMinProducerSlot + (producerDistance * uint32(common.DefaultConfig.ProducerRepetitions))
		result = common.BlockTimeStamp(nextBlockSlot).ToTimePoint()

	}
	return &result
}

func (impl *ProducerPluginImpl) CalculatePendingBlockTime() common.TimePoint {
	chain := Chain.GetControllerInstance()
	now := common.Now()
	var base common.TimePoint
	if now > chain.HeadBlockTime() {
		base = now
	} else {
		base = chain.HeadBlockTime()
	}
	minTimeToNextBlock := common.DefaultConfig.BlockIntervalUs - (int64(base.TimeSinceEpoch()) % common.DefaultConfig.BlockIntervalUs)
	blockTime := base.AddUs(common.Microseconds(minTimeToNextBlock))

	if blockTime.Sub(now) < common.Microseconds(common.DefaultConfig.BlockIntervalUs/10) { // we must sleep for at least 50ms
		blockTime = blockTime.AddUs(common.Microseconds(common.DefaultConfig.BlockIntervalUs))
	}

	return blockTime
}

func (impl *ProducerPluginImpl) StartBlock() (EnumStartBlockRusult, bool) {
	chain := Chain.GetControllerInstance()

	if chain.GetReadMode() == Chain.READONLY {
		return EnumStartBlockRusult(waiting), false
	}

	hbs := chain.HeadBlockState()

	//Schedule for the next second's tick regardless of chain state
	// If we would wait less than 50ms (1/10 of block_interval), wait for the whole block interval.
	now := common.Now()
	blockTime := impl.CalculatePendingBlockTime()

	impl.PendingBlockMode = EnumPendingBlockMode(producing)

	// Not our turn
	lastBlock := uint32(common.NewBlockTimeStamp(blockTime))%uint32(common.DefaultConfig.ProducerRepetitions) == uint32(common.DefaultConfig.ProducerRepetitions)-1
	scheduleProducer := hbs.GetScheduledProducer(common.NewBlockTimeStamp(blockTime))
	currentWatermark, hasCurrentWatermark := impl.ProducerWatermarks[scheduleProducer.AccountName]
	_, hasSignatureProvider := impl.SignatureProviders[scheduleProducer.BlockSigningKey]
	irreversibleBlockAge := impl.GetIrreversibleBlockAge()

	// If the next block production opportunity is in the present or future, we're synced.
	if !impl.ProductionEnabled {
		impl.PendingBlockMode = EnumPendingBlockMode(speculating)
	} else if _, has := impl.Producers[scheduleProducer.AccountName]; !has {
		impl.PendingBlockMode = EnumPendingBlockMode(speculating)
	} else if !hasSignatureProvider {
		impl.PendingBlockMode = EnumPendingBlockMode(speculating)
		//elog("Not producing block because I don't have the private key for ${scheduled_key}", ("scheduled_key", scheduled_producer.block_signing_key));
	} else if impl.ProductionPaused {
		//elog("Not producing block because production is explicitly paused");
		impl.PendingBlockMode = EnumPendingBlockMode(speculating)
	} else if impl.MaxIrreversibleBlockAgeUs >= 0 && irreversibleBlockAge >= impl.MaxIrreversibleBlockAgeUs {
		//elog("Not producing block because the irreversible block is too old [age:${age}s, max:${max}s]", ("age", irreversible_block_age.count() / 1'000'000)( "max", _max_irreversible_block_age_us.count() / 1'000'000 ));
		impl.PendingBlockMode = EnumPendingBlockMode(speculating)
	}

	if impl.PendingBlockMode == EnumPendingBlockMode(producing) {
		if hasCurrentWatermark {
			if currentWatermark >= hbs.BlockNum+1 {
				/*
									elog("Not producing block because \"${producer}\" signed a BFT confirmation OR block at a higher block number (${watermark}) than the current fork's head (${head_block_num})",
					                ("producer", scheduled_producer.producer_name)
					                ("watermark", currrent_watermark_itr->second)
					                ("head_block_num", hbs->block_num));
				*/
				impl.PendingBlockMode = EnumPendingBlockMode(speculating)
			}

		}
	}

	if impl.PendingBlockMode == EnumPendingBlockMode(speculating) {
		headBlockAge := now.Sub(chain.HeadBlockTime())
		if headBlockAge > common.Seconds(5) {
			return EnumStartBlockRusult(waiting), lastBlock
		}
	}

	var blocksToConfirm uint16 = 0

	if impl.PendingBlockMode == EnumPendingBlockMode(producing) {
		// determine how many blocks this producer can confirm
		// 1) if it is not a producer from this node, assume no confirmations (we will discard this block anyway)
		// 2) if it is a producer on this node that has never produced, the conservative approach is to assume no
		//    confirmations to make sure we don't double sign after a crash TODO: make these watermarks durable?
		// 3) if it is a producer on this node where this node knows the last block it produced, safely set it -UNLESS-
		// 4) the producer on this node's last watermark is higher (meaning on a different fork)
		if hasCurrentWatermark {
			if currentWatermark < hbs.BlockNum {
				if hbs.BlockNum-currentWatermark >= 0xffff {
					blocksToConfirm = 0xffff
				} else {
					blocksToConfirm = uint16(hbs.BlockNum - currentWatermark)
				}
			}
		}
	}

	chain.AbortBlock()
	chain.StartBlock(common.NewBlockTimeStamp(blockTime), blocksToConfirm)

	pbs := chain.PendingBlockState()

	if pbs != nil {

		if impl.PendingBlockMode == EnumPendingBlockMode(producing) && pbs.BlockSigningKey != scheduleProducer.BlockSigningKey {
			//C++ elog("Block Signing Key is not expected value, reverting to speculative mode! [expected: \"${expected}\", actual: \"${actual\"", ("expected", scheduled_producer.block_signing_key)("actual", pbs->block_signing_key));
			impl.PendingBlockMode = EnumPendingBlockMode(speculating)
		}

		// attempt to play persisted transactions first
		isExhausted := false

		// remove all persisted transactions that have now expired
		for byTrxId, byExpire := range impl.PersistentTransactions {
			if byExpire <= pbs.Header.Timestamp.ToTimePoint() {
				delete(impl.PersistentTransactions, byTrxId)
			}
		}

		origPendingTxnSize := len(impl.PendingIncomingTransactions)

		if len(impl.PersistentTransactions) > 0 || impl.PendingBlockMode == EnumPendingBlockMode(producing) {
			unappliedTrxs := chain.GetUnappliedTransactions()

			if len(impl.PersistentTransactions) > 0 {
				for i, trx := range unappliedTrxs {
					if _, has := impl.PersistentTransactions[trx.ID]; has {
						// this is a persisted transaction, push it into the block (even if we are speculating) with
						// no deadline as it has already passed the subjective deadlines once and we want to represent
						// the state of the chain including this transaction
						err := chain.PushTransaction(trx, common.MaxTimePoint())
						if err != nil {
							return EnumStartBlockRusult(failed), lastBlock
						}
					}

					// remove it from further consideration as it is applied
					unappliedTrxs[i] = nil
				}
			}

			if impl.PendingBlockMode == EnumPendingBlockMode(producing) {
				for _, trx := range unappliedTrxs {
					if blockTime <= common.Now() {
						isExhausted = true
					}
					if isExhausted {
						break
					}

					if trx == nil {
						// nulled in the loop above, skip it
						continue
					}

					if trx.PackedTrx.Expiration().ToTimePoint() < pbs.Header.Timestamp.ToTimePoint() {
						// expired, drop it
						chain.DropUnappliedTransaction(trx)
						continue
					}

					deadline := common.Now().AddUs(common.Microseconds(impl.MaxTransactionTimeMs))
					deadlineIsSubjective := false
					if impl.MaxTransactionTimeMs < 0 || impl.PendingBlockMode == EnumPendingBlockMode(producing) && blockTime < deadline {
						deadlineIsSubjective = true
						deadline = blockTime
					}

					trace := chain.PushTransaction(trx, deadline)
					if trace.Except != nil {
						if failureIsSubjective(trace.Except, deadlineIsSubjective) {
							isExhausted = true
						} else {
							// this failed our configured maximum transaction time, we don't want to replay it
							chain.DropUnappliedTransaction(trx)
						}
					}

					//TODO catch exception
				}
			}
		} ///unapplied transactions

		if impl.PendingBlockMode == EnumPendingBlockMode(producing) {
			for byTrxId, byExpire := range impl.BlacklistedTransactions {
				if byExpire <= common.Now() {
					delete(impl.BlacklistedTransactions, byTrxId)
				}
			}

			scheduledTrxs := chain.GetScheduledTransactions()

			for _, trx := range scheduledTrxs {
				if blockTime <= common.Now() {
					isExhausted = true
				}
				if isExhausted {
					break
				}

				// configurable ratio of incoming txns vs deferred txns
				for impl.IncomingTrxWeight >= 1.0 && origPendingTxnSize > 0 && len(impl.PendingIncomingTransactions) > 0 {
					e := impl.PendingIncomingTransactions[0]
					impl.PendingIncomingTransactions = impl.PendingIncomingTransactions[1:]
					origPendingTxnSize--
					impl.IncomingTrxWeight -= 1.0
					impl.OnIncomingTransactionAsync(e.packedTransaction, e.persistUntilExpired, e.next)
				}

				if blockTime <= common.Now() {
					isExhausted = true
					break
				}

				if _, has := impl.BlacklistedTransactions[trx]; has {
					continue
				}

				deadline := common.Now().AddUs(common.Microseconds(impl.MaxTransactionTimeMs))
				deadlineIsSubjective := false
				if impl.MaxTransactionTimeMs < 0 || impl.PendingBlockMode == EnumPendingBlockMode(producing) && blockTime < deadline {
					deadlineIsSubjective = true
					deadline = blockTime
				}

				trace := chain.PushScheduledTransaction(trx, deadline)
				if trace.Except != nil {
					if failureIsSubjective(trace.Except, deadlineIsSubjective) {
						isExhausted = true
					} else {
						expiration := common.Now().AddUs(common.Seconds(0) /*TODO chain.get_global_properties().configuration.deferred_trx_expiration_window*/)
						// this failed our configured maximum transaction time, we don't want to replay it add it to a blacklist
						impl.BlacklistedTransactions[trx] = expiration
					}
				}

				//TODO catch exception

				impl.IncomingTrxWeight += impl.IncomingDeferRadio
				if origPendingTxnSize <= 0 {
					impl.IncomingTrxWeight = 0.0
				}
			}
		} ///scheduled transactions

		if isExhausted || blockTime <= common.Now() {
			return EnumStartBlockRusult(exhausted), lastBlock
		} else {
			// attempt to apply any pending incoming transactions
			impl.IncomingTrxWeight = 0.0
			if origPendingTxnSize > 0 && len(impl.PendingIncomingTransactions) > 0 {
				e := impl.PendingIncomingTransactions[0]
				impl.PendingIncomingTransactions = impl.PendingIncomingTransactions[1:]
				origPendingTxnSize--
				impl.OnIncomingTransactionAsync(e.packedTransaction, e.persistUntilExpired, e.next)
				if blockTime <= common.Now() {
					return EnumStartBlockRusult(exhausted), lastBlock
				}
			}
			return EnumStartBlockRusult(succeeded), lastBlock
		}
	}
	return EnumStartBlockRusult(failed), lastBlock
}

func (impl *ProducerPluginImpl) ScheduleProductionLoop() {
	chain := Chain.GetControllerInstance()
	impl.Timer.Cancel()

	result, lastBlock := impl.StartBlock()

	if result == EnumStartBlockRusult(failed) {
		//elog("Failed to start a pending block, will try again later");
		impl.Timer.ExpiresFromNow(common.Microseconds(common.DefaultConfig.BlockIntervalUs / 10))

		// we failed to start a block, so try again later?
		impl.timerCorelationId++
		cid := impl.timerCorelationId
		impl.Timer.AsyncWait(func() {
			if impl != nil && cid == impl.timerCorelationId {
				impl.ScheduleProductionLoop()
			}
		})

	} else if result == EnumStartBlockRusult(waiting) {
		if len(impl.Producers) > 0 && !impl.ProductionDisabledByPolicy() {
			log.Debug("Waiting till another block is received and scheduling Speculative/Production Change")
			impl.ScheduleDelayedProductionLoop(common.NewBlockTimeStamp(impl.CalculatePendingBlockTime()))
		} else {
			log.Debug("Waiting till another block is received")
			// nothing to do until more blocks arrive
		}

	} else if impl.PendingBlockMode == EnumPendingBlockMode(producing) {
		// we succeeded but block may be exhausted
		if result == EnumStartBlockRusult(succeeded) {
			// ship this block off no later than its deadline
			if chain.PendingBlockState() == nil {
				panic("producing without pending_block_state, start_block succeeded")
			}
			epoch := chain.PendingBlockTime().TimeSinceEpoch()
			if lastBlock {
				epoch += common.Microseconds(impl.LastBlockTimeOffsetUs)
			} else {
				epoch += common.Microseconds(impl.ProduceTimeOffsetUs)
			}
			impl.Timer.ExpiresAt(epoch)
			log.Debug(fmt.Sprintf("Scheduling Block Production on Normal Block #%dfor %s", chain.PendingBlockState().BlockNum, chain.PendingBlockTime()))
		} else {
			expectTime := chain.PendingBlockTime().SubUs(common.Microseconds(common.DefaultConfig.BlockIntervalUs))
			// ship this block off up to 1 block time earlier or immediately
			if common.Now() >= expectTime {
				impl.Timer.ExpiresFromNow(0)
			} else {
				impl.Timer.ExpiresAt(expectTime.TimeSinceEpoch())
			}
			log.Debug("Scheduling Block Production on Exhausted Block #%d immediately", chain.PendingBlockState().BlockNum)
		}

		impl.timerCorelationId++
		cid := impl.timerCorelationId
		impl.Timer.AsyncWait(func() {
			if impl != nil && cid == impl.timerCorelationId {
				impl.MaybeProduceBlock()
				log.Debug("Producing Block #${num} returned: ${res}")
				//fc_dlog(_log, "Producing Block #${num} returned: ${res}", ("num", chain.pending_block_state()->block_num)("res", res) );
			}
		})

	} else if impl.PendingBlockMode == EnumPendingBlockMode(speculating) && len(impl.Producers) > 0 && !impl.ProductionDisabledByPolicy() {
		//fc_dlog(_log, "Specualtive Block Created; Scheduling Speculative/Production Change");
		pbs := chain.PendingBlockState()
		impl.ScheduleDelayedProductionLoop(pbs.Header.Timestamp)

	} else {
		//fc_dlog(_log, "Speculative Block Created");
	}
}

func (impl *ProducerPluginImpl) ScheduleDelayedProductionLoop(currentBlockTime common.BlockTimeStamp) {
	var wakeUpTime *common.TimePoint
	for p := range impl.Producers {
		nextProducerBlockTime := impl.CalculateNextBlockTime(p, currentBlockTime)
		if nextProducerBlockTime != nil {
			producerWakeupTime := nextProducerBlockTime.SubUs(common.Microseconds(common.DefaultConfig.BlockIntervalUs))
			if wakeUpTime != nil {
				if *wakeUpTime > producerWakeupTime {
					*wakeUpTime = producerWakeupTime
				}
			} else {
				wakeUpTime = &producerWakeupTime
			}
		}
	}

	if wakeUpTime != nil {
		//fc_dlog(_log, "Scheduling Speculative/Production Change at ${time}", ("time", wake_up_time));
		impl.Timer.ExpiresAt(wakeUpTime.TimeSinceEpoch())

		impl.timerCorelationId++
		cid := impl.timerCorelationId
		impl.Timer.AsyncWait(func() {
			if impl != nil && cid == impl.timerCorelationId {
				impl.ScheduleProductionLoop()
			}
		})
	} else {
		//fc_dlog(_log, "Speculative Block Created; Not Scheduling Speculative/Production, no local producers had valid wake up times");
	}

}

func (impl *ProducerPluginImpl) MaybeProduceBlock() (res bool) {
	defer func() {
		impl.ScheduleProductionLoop()
	}()

	try.Try(func() {
		impl.ProduceBlock()
		res = true
	}).Catch(func(e GuardExceptions) {
		res = false
	}).Catch(func(e Exception) {
		//fc_dlog(_log, "Aborting block due to produce_block error");
		chain := Chain.GetControllerInstance()
		chain.AbortBlock()
		res = false
	}).End()

	return
}

func (impl *ProducerPluginImpl) ProduceBlock() {
	if impl.PendingBlockMode != EnumPendingBlockMode(producing) {
		panic(ErrProducerFail)
	}
	chain := Chain.GetControllerInstance()
	pbs := chain.PendingBlockState()
	if pbs == nil {
		panic(ErrMissingPendingBlockState)
	}

	signatureProvider := impl.SignatureProviders[pbs.BlockSigningKey]
	if signatureProvider == nil {
		panic(ErrProducerPriKeyNotFound)
	}

	chain.FinalizeBlock()
	chain.SignBlock(func(d crypto.Sha256) ecc.Signature {
		defer makeDebugTimeLogger()
		return signatureProvider(d)
	})

	chain.CommitBlock(true)

	newBs := chain.HeadBlockState()
	impl.ProducerWatermarks[newBs.Header.Producer] = chain.HeadBlockNum()

	fmt.Printf("Produced block %s...#%d @ %s signed by %s [trxs: %d, lib: %d, confirmed: %d]\n",
		newBs.ID.String()[0:16], newBs.BlockNum, newBs.Header.Timestamp, common.S(uint64(newBs.Header.Producer)),
		len(newBs.SignedBlock.Transactions), chain.LastIrreversibleBlockNum(), newBs.Header.Confirmed)
}