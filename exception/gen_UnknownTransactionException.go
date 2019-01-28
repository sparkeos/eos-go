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

var UnknownTransactionExceptionName = reflect.TypeOf(UnknownTransactionException{}).Name()

type UnknownTransactionException struct {
	_MiscException
	Elog log.Messages
}

func NewUnknownTransactionException(parent _MiscException, message log.Message) *UnknownTransactionException {
	return &UnknownTransactionException{parent, log.Messages{message}}
}

func (e UnknownTransactionException) Code() int64 {
	return 3100003
}

func (e UnknownTransactionException) Name() string {
	return UnknownTransactionExceptionName
}

func (e UnknownTransactionException) What() string {
	return "Unknown transaction"
}

func (e *UnknownTransactionException) AppendLog(l log.Message) {
	e.Elog = append(e.Elog, l)
}

func (e UnknownTransactionException) GetLog() log.Messages {
	return e.Elog
}

func (e UnknownTransactionException) TopMessage() string {
	for _, l := range e.Elog {
		if msg := l.GetMessage(); len(msg) > 0 {
			return msg
		}
	}
	return e.String()
}

func (e UnknownTransactionException) DetailMessage() string {
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

func (e UnknownTransactionException) String() string {
	return e.DetailMessage()
}

func (e UnknownTransactionException) MarshalJSON() ([]byte, error) {
	type Exception struct {
		Code int64  `json:"code"`
		Name string `json:"name"`
		What string `json:"what"`
	}

	except := Exception{
		Code: 3100003,
		Name: UnknownTransactionExceptionName,
		What: "Unknown transaction",
	}

	return json.Marshal(except)
}

func (e UnknownTransactionException) Callback(f interface{}) bool {
	switch callback := f.(type) {
	case func(*UnknownTransactionException):
		callback(&e)
		return true
	case func(UnknownTransactionException):
		callback(e)
		return true
	default:
		return false
	}
}
