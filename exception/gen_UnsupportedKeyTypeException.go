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

var UnsupportedKeyTypeExceptionName = reflect.TypeOf(UnsupportedKeyTypeException{}).Name()

type UnsupportedKeyTypeException struct {
	_WalletException
	Elog log.Messages
}

func NewUnsupportedKeyTypeException(parent _WalletException, message log.Message) *UnsupportedKeyTypeException {
	return &UnsupportedKeyTypeException{parent, log.Messages{message}}
}

func (e UnsupportedKeyTypeException) Code() int64 {
	return 3120010
}

func (e UnsupportedKeyTypeException) Name() string {
	return UnsupportedKeyTypeExceptionName
}

func (e UnsupportedKeyTypeException) What() string {
	return "Unsupported key type"
}

func (e *UnsupportedKeyTypeException) AppendLog(l log.Message) {
	e.Elog = append(e.Elog, l)
}

func (e UnsupportedKeyTypeException) GetLog() log.Messages {
	return e.Elog
}

func (e UnsupportedKeyTypeException) TopMessage() string {
	for _, l := range e.Elog {
		if msg := l.GetMessage(); msg != "" {
			return msg
		}
	}
	return e.String()
}

func (e UnsupportedKeyTypeException) DetailMessage() string {
	var buffer bytes.Buffer
	buffer.WriteString(strconv.Itoa(int(e.Code())))
	buffer.WriteString(" ")
	buffer.WriteString(e.Name())
	buffer.WriteString(": ")
	buffer.WriteString(e.What())
	buffer.WriteString("\n")
	for _, l := range e.Elog {
		buffer.WriteString("[")
		buffer.WriteString(l.GetMessage())
		buffer.WriteString("]")
		buffer.WriteString("\n")
		buffer.WriteString(l.GetContext().String())
		buffer.WriteString("\n")
	}
	return buffer.String()
}

func (e UnsupportedKeyTypeException) String() string {
	return e.DetailMessage()
}

func (e UnsupportedKeyTypeException) MarshalJSON() ([]byte, error) {
	type Exception struct {
		Code int64  `json:"code"`
		Name string `json:"name"`
		What string `json:"what"`
	}

	except := Exception{
		Code: 3120010,
		Name: UnsupportedKeyTypeExceptionName,
		What: "Unsupported key type",
	}

	return json.Marshal(except)
}

func (e UnsupportedKeyTypeException) Callback(f interface{}) bool {
	switch callback := f.(type) {
	case func(*UnsupportedKeyTypeException):
		callback(&e)
		return true
	case func(UnsupportedKeyTypeException):
		callback(e)
		return true
	default:
		return false
	}
}
