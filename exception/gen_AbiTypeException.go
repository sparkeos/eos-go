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

var AbiTypeExceptionName = reflect.TypeOf(AbiTypeException{}).Name()

type AbiTypeException struct {
	_ChainException
	Elog log.Messages
}

func NewAbiTypeException(parent _ChainException, message log.Message) *AbiTypeException {
	return &AbiTypeException{parent, log.Messages{message}}
}

func (e AbiTypeException) Code() int64 {
	return 3010007
}

func (e AbiTypeException) Name() string {
	return AbiTypeExceptionName
}

func (e AbiTypeException) What() string {
	return "Invalid ABI"
}

func (e *AbiTypeException) AppendLog(l log.Message) {
	e.Elog = append(e.Elog, l)
}

func (e AbiTypeException) GetLog() log.Messages {
	return e.Elog
}

func (e AbiTypeException) TopMessage() string {
	for _, l := range e.Elog {
		if msg := l.GetMessage(); len(msg) > 0 {
			return msg
		}
	}
	return e.String()
}

func (e AbiTypeException) DetailMessage() string {
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

func (e AbiTypeException) String() string {
	return e.DetailMessage()
}

func (e AbiTypeException) MarshalJSON() ([]byte, error) {
	type Exception struct {
		Code int64  `json:"code"`
		Name string `json:"name"`
		What string `json:"what"`
	}

	except := Exception{
		Code: 3010007,
		Name: AbiTypeExceptionName,
		What: "Invalid ABI",
	}

	return json.Marshal(except)
}

func (e AbiTypeException) Callback(f interface{}) bool {
	switch callback := f.(type) {
	case func(*AbiTypeException):
		callback(&e)
		return true
	case func(AbiTypeException):
		callback(e)
		return true
	default:
		return false
	}
}
