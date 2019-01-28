// Code generated by gotemplate. DO NOT EDIT.

package exception

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strconv"

	"github.com/eosspark/eos-go/log"
)

// template type Exception(PARENT,CODE,WHAT)

var DeferredTxDuplicateName = reflect.TypeOf(DeferredTxDuplicate{}).Name()

type DeferredTxDuplicate struct {
	_TransactionException
	Elog log.Messages
}

func NewDeferredTxDuplicate(parent _TransactionException, message log.Message) *DeferredTxDuplicate {
	return &DeferredTxDuplicate{parent, log.Messages{message}}
}

func (e DeferredTxDuplicate) Code() int64 {
	return 3040009
}

func (e DeferredTxDuplicate) Name() string {
	return DeferredTxDuplicateName
}

func (e DeferredTxDuplicate) What() string {
	return "Duplicate deferred transaction"
}

func (e *DeferredTxDuplicate) AppendLog(l log.Message) {
	e.Elog = append(e.Elog, l)
}

func (e DeferredTxDuplicate) GetLog() log.Messages {
	return e.Elog
}

func (e DeferredTxDuplicate) TopMessage() string {
	for _, l := range e.Elog {
		if msg := l.GetMessage(); len(msg) > 0 {
			return msg
		}
	}
	return e.String()
}

func (e DeferredTxDuplicate) DetailMessage() string {
	var buffer bytes.Buffer
	buffer.WriteString(strconv.Itoa(int(e.Code())))
	buffer.WriteByte(' ')
	buffer.WriteString(e.Name())
	buffer.Write([]byte{':', ' '})
	buffer.WriteString(e.What())
	buffer.WriteByte('\n')
	for _, l := range e.Elog {
		buffer.WriteByte('[')
		buffer.WriteString(l.GetMessage())
		buffer.Write([]byte{']', ' '})
		buffer.WriteString(l.GetContext().String())
		buffer.WriteByte('\n')
	}
	return buffer.String()
}

func (e DeferredTxDuplicate) String() string {
	return e.DetailMessage()
}

func (e DeferredTxDuplicate) MarshalJSON() ([]byte, error) {
	type Exception struct {
		Code int64  `json:"code"`
		Name string `json:"name"`
		What string `json:"what"`
	}

	except := Exception{
		Code: 3040009,
		Name: DeferredTxDuplicateName,
		What: "Duplicate deferred transaction",
	}

	return json.Marshal(except)
}

func (e DeferredTxDuplicate) Callback(f interface{}) bool {
	switch callback := f.(type) {
	case func(*DeferredTxDuplicate):
		callback(&e)
		return true
	case func(DeferredTxDuplicate):
		callback(e)
		return true
	default:
		return false
	}
}
