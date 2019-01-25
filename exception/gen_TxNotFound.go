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

var TxNotFoundName = reflect.TypeOf(TxNotFound{}).Name()

type TxNotFound struct {
	_TransactionException
	Elog log.Messages
}

func NewTxNotFound(parent _TransactionException, message log.Message) *TxNotFound {
	return &TxNotFound{parent, log.Messages{message}}
}

func (e TxNotFound) Code() int64 {
	return 3040011
}

func (e TxNotFound) Name() string {
	return TxNotFoundName
}

func (e TxNotFound) What() string {
	return "The transaction can not be found"
}

func (e *TxNotFound) AppendLog(l log.Message) {
	e.Elog = append(e.Elog, l)
}

func (e TxNotFound) GetLog() log.Messages {
	return e.Elog
}

func (e TxNotFound) TopMessage() string {
	for _, l := range e.Elog {
		if msg := l.GetMessage(); len(msg) > 0 {
			return msg
		}
	}
	return e.String()
}

func (e TxNotFound) DetailMessage() string {
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

func (e TxNotFound) String() string {
	return e.DetailMessage()
}

func (e TxNotFound) MarshalJSON() ([]byte, error) {
	type Exception struct {
		Code int64  `json:"code"`
		Name string `json:"name"`
		What string `json:"what"`
	}

	except := Exception{
		Code: 3040011,
		Name: TxNotFoundName,
		What: "The transaction can not be found",
	}

	return json.Marshal(except)
}

func (e TxNotFound) Callback(f interface{}) bool {
	switch callback := f.(type) {
	case func(*TxNotFound):
		callback(&e)
		return true
	case func(TxNotFound):
		callback(e)
		return true
	default:
		return false
	}
}
