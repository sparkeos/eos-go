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

var PackedTransactionTypeExceptionName = reflect.TypeOf(PackedTransactionTypeException{}).Name()

type PackedTransactionTypeException struct {
	_ChainException
	Elog log.Messages
}

func NewPackedTransactionTypeException(parent _ChainException, message log.Message) *PackedTransactionTypeException {
	return &PackedTransactionTypeException{parent, log.Messages{message}}
}

func (e PackedTransactionTypeException) Code() int64 {
	return 3010010
}

func (e PackedTransactionTypeException) Name() string {
	return PackedTransactionTypeExceptionName
}

func (e PackedTransactionTypeException) What() string {
	return "Invalid packed transaction"
}

func (e *PackedTransactionTypeException) AppendLog(l log.Message) {
	e.Elog = append(e.Elog, l)
}

func (e PackedTransactionTypeException) GetLog() log.Messages {
	return e.Elog
}

func (e PackedTransactionTypeException) TopMessage() string {
	for _, l := range e.Elog {
		if msg := l.GetMessage(); len(msg) > 0 {
			return msg
		}
	}
	return e.String()
}

func (e PackedTransactionTypeException) DetailMessage() string {
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

func (e PackedTransactionTypeException) String() string {
	return e.DetailMessage()
}

func (e PackedTransactionTypeException) MarshalJSON() ([]byte, error) {
	type Exception struct {
		Code int64  `json:"code"`
		Name string `json:"name"`
		What string `json:"what"`
	}

	except := Exception{
		Code: 3010010,
		Name: PackedTransactionTypeExceptionName,
		What: "Invalid packed transaction",
	}

	return json.Marshal(except)
}

func (e PackedTransactionTypeException) Callback(f interface{}) bool {
	switch callback := f.(type) {
	case func(*PackedTransactionTypeException):
		callback(&e)
		return true
	case func(PackedTransactionTypeException):
		callback(e)
		return true
	default:
		return false
	}
}
