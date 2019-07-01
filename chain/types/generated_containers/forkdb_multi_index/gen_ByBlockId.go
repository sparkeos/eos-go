// Code generated by gotemplate. DO NOT EDIT.

package forkdb_multi_index

import (
	"github.com/eosspark/eos-go/common"
	"github.com/eosspark/eos-go/libraries/container"
	"github.com/eosspark/eos-go/libraries/multiindex"
)

// template type HashedUniqueIndex(FinalIndex,FinalNode,SuperIndex,SuperNode,Value,Hash,KeyFunc)

type ByBlockId struct {
	super *ByPrev                               // index on the HashedUniqueIndex, IndexBase is the last super index
	final *MultiIndex                           // index under the HashedUniqueIndex, MultiIndex is the final index
	inner map[common.BlockIdType]*ByBlockIdNode // use hashmap to safe HashedUniqueIndex's k/v(HashedUniqueIndexNode)
}

func (i *ByBlockId) init(final *MultiIndex) {
	i.final = final
	i.inner = map[common.BlockIdType]*ByBlockIdNode{}
	i.super = &ByPrev{}
	i.super.init(final)
}

/*generic class*/

/*generic class*/

type ByBlockIdNode struct {
	super *ByPrevNode        // index-node on the HashedUniqueIndexNode, IndexBaseNode is the last super node
	final *MultiIndexNode    // index-node under the HashedUniqueIndexNode, MultiIndexNode is the final index
	hash  common.BlockIdType // k of hashmap
}

/*generic class*/

/*generic class*/

func (i *ByBlockId) GetSuperIndex() interface{} { return i.super }
func (i *ByBlockId) GetFinalIndex() interface{} { return i.final }

func (n *ByBlockIdNode) GetSuperNode() interface{} { return n.super }
func (n *ByBlockIdNode) GetFinalNode() interface{} { return n.final }

func (n *ByBlockIdNode) value() *BlockStatePtr {
	return n.super.value()
}

func (i *ByBlockId) Size() int {
	return len(i.inner)
}

func (i *ByBlockId) Empty() bool {
	return len(i.inner) == 0
}

func (i *ByBlockId) clear() {
	i.inner = map[common.BlockIdType]*ByBlockIdNode{}
	i.super.clear()
}

func (i *ByBlockId) Insert(v BlockStatePtr) (IteratorByBlockId, bool) {
	fn, res := i.final.insert(v)
	if res {
		return i.makeIterator(fn), true
	}
	return i.End(), false
}

func (i *ByBlockId) insert(v BlockStatePtr, fn *MultiIndexNode) (*ByBlockIdNode, bool) {
	hash := ByBlockIdFunc(v)
	node := ByBlockIdNode{hash: hash}
	if _, ok := i.inner[hash]; ok {
		container.Logger.Warn("#hash index insert failed")
		return nil, false
	}
	i.inner[hash] = &node
	sn, res := i.super.insert(v, fn)
	if res {
		node.final = fn
		node.super = sn
		return &node, true
	}
	delete(i.inner, hash)
	return nil, false
}

func (i *ByBlockId) Find(k common.BlockIdType) (IteratorByBlockId, bool) {
	node, res := i.inner[k]
	if res {
		return IteratorByBlockId{i, node, betweenByBlockId}, true
	}
	return i.End(), false
}

func (i *ByBlockId) Each(f func(key common.BlockIdType, obj BlockStatePtr)) {
	for k, v := range i.inner {
		f(k, *v.value())
	}
}

func (i *ByBlockId) Erase(iter IteratorByBlockId) {
	i.final.erase(iter.node.final)
}

func (i *ByBlockId) erase(n *ByBlockIdNode) {
	delete(i.inner, n.hash)
	i.super.erase(n.super)
}

func (i *ByBlockId) erase_(iter multiindex.IteratorType) {
	if itr, ok := iter.(IteratorByBlockId); ok {
		i.Erase(itr)
	} else {
		i.super.erase_(iter)
	}
}

func (i *ByBlockId) Modify(iter IteratorByBlockId, mod func(*BlockStatePtr)) bool {
	if _, b := i.final.modify(mod, iter.node.final); b {
		return true
	}
	return false
}

func (i *ByBlockId) modify(n *ByBlockIdNode) (*ByBlockIdNode, bool) {
	delete(i.inner, n.hash)

	hash := ByBlockIdFunc(*n.value())
	if _, exist := i.inner[hash]; exist {
		container.Logger.Warn("#hash index modify failed")
		i.super.erase(n.super)
		return nil, false
	}

	i.inner[hash] = n

	if sn, res := i.super.modify(n.super); !res {
		delete(i.inner, hash)
		return nil, false
	} else {
		n.super = sn
	}

	return n, true
}

func (i *ByBlockId) modify_(iter multiindex.IteratorType, mod func(*BlockStatePtr)) bool {
	if itr, ok := iter.(IteratorByBlockId); ok {
		return i.Modify(itr, mod)
	} else {
		return i.super.modify_(iter, mod)
	}
}

func (i *ByBlockId) Values() []BlockStatePtr {
	vs := make([]BlockStatePtr, 0, i.Size())
	i.Each(func(key common.BlockIdType, obj BlockStatePtr) {
		vs = append(vs, obj)
	})
	return vs
}

type IteratorByBlockId struct {
	index    *ByBlockId
	node     *ByBlockIdNode
	position posByBlockId
}

type posByBlockId byte

const (
	//begin   = 0
	betweenByBlockId = 1
	endByBlockId     = 2
)

func (i *ByBlockId) makeIterator(fn *MultiIndexNode) IteratorByBlockId {
	node := fn.GetSuperNode()
	for {
		if node == nil {
			panic("Wrong index node type!")

		} else if n, ok := node.(*ByBlockIdNode); ok {
			return IteratorByBlockId{i, n, betweenByBlockId}
		} else {
			node = node.(multiindex.NodeType).GetSuperNode()
		}
	}
}

func (i *ByBlockId) End() IteratorByBlockId {
	return IteratorByBlockId{i, nil, endByBlockId}
}

func (iter IteratorByBlockId) Value() (v BlockStatePtr) {
	if iter.position == betweenByBlockId {
		return *iter.node.value()
	}
	return
}

func (iter IteratorByBlockId) HasNext() bool {
	container.Logger.Warn("hashed index iterator is unmoveable")
	return false
}

func (iter IteratorByBlockId) IsEnd() bool {
	return iter.position == endByBlockId
}