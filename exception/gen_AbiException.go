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

var AbiExceptionName = reflect.TypeOf(AbiException{}).Name()

type AbiException struct {
	_AbiException
	Elog log.Messages
}

func NewAbiException(parent _AbiException, message log.Message) *AbiException {
	return &AbiException{parent, log.Messages{message}}
}

func (e AbiException) Code() int64 {
	return 3150000
}

func (e AbiException) Name() string {
	return AbiExceptionName
}

func (e AbiException) What() string {
	return "ABI exception"
}

func (e *AbiException) AppendLog(l log.Message) {
	e.Elog = append(e.Elog, l)
}

func (e AbiException) GetLog() log.Messages {
	return e.Elog
}

func (e AbiException) TopMessage() string {
	for _, l := range e.Elog {
		if msg := l.GetMessage(); msg != "" {
			return msg
		}
	}
	return e.String()
}

func (e AbiException) DetailMessage() string {
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

func (e AbiException) String() string {
	return e.DetailMessage()
}

func (e AbiException) MarshalJSON() ([]byte, error) {
	type Exception struct {
		Code int64  `json:"code"`
		Name string `json:"name"`
		What string `json:"what"`
	}

	except := Exception{
		Code: 3150000,
		Name: AbiExceptionName,
		What: "ABI exception",
	}

	return json.Marshal(except)
}

func (e AbiException) Callback(f interface{}) bool {
	switch callback := f.(type) {
	case func(*AbiException):
		callback(&e)
		return true
	case func(AbiException):
		callback(e)
		return true
	default:
		return false
	}
}