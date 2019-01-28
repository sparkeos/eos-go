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

var InvalidTableIteratorName = reflect.TypeOf(InvalidTableIterator{}).Name()

type InvalidTableIterator struct {
	_ContractException
	Elog log.Messages
}

func NewInvalidTableIterator(parent _ContractException, message log.Message) *InvalidTableIterator {
	return &InvalidTableIterator{parent, log.Messages{message}}
}

func (e InvalidTableIterator) Code() int64 {
	return 3160003
}

func (e InvalidTableIterator) Name() string {
	return InvalidTableIteratorName
}

func (e InvalidTableIterator) What() string {
	return "Invalid table iterator"
}

func (e *InvalidTableIterator) AppendLog(l log.Message) {
	e.Elog = append(e.Elog, l)
}

func (e InvalidTableIterator) GetLog() log.Messages {
	return e.Elog
}

func (e InvalidTableIterator) TopMessage() string {
	for _, l := range e.Elog {
		if msg := l.GetMessage(); len(msg) > 0 {
			return msg
		}
	}
	return e.String()
}

func (e InvalidTableIterator) DetailMessage() string {
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

func (e InvalidTableIterator) String() string {
	return e.DetailMessage()
}

func (e InvalidTableIterator) MarshalJSON() ([]byte, error) {
	type Exception struct {
		Code int64  `json:"code"`
		Name string `json:"name"`
		What string `json:"what"`
	}

	except := Exception{
		Code: 3160003,
		Name: InvalidTableIteratorName,
		What: "Invalid table iterator",
	}

	return json.Marshal(except)
}

func (e InvalidTableIterator) Callback(f interface{}) bool {
	switch callback := f.(type) {
	case func(*InvalidTableIterator):
		callback(&e)
		return true
	case func(InvalidTableIterator):
		callback(e)
		return true
	default:
		return false
	}
}
