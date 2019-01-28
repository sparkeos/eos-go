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

var UnsupportedAbiVersionExceptionName = reflect.TypeOf(UnsupportedAbiVersionException{}).Name()

type UnsupportedAbiVersionException struct {
	_AbiException
	Elog log.Messages
}

func NewUnsupportedAbiVersionException(parent _AbiException, message log.Message) *UnsupportedAbiVersionException {
	return &UnsupportedAbiVersionException{parent, log.Messages{message}}
}

func (e UnsupportedAbiVersionException) Code() int64 {
	return 3150016
}

func (e UnsupportedAbiVersionException) Name() string {
	return UnsupportedAbiVersionExceptionName
}

func (e UnsupportedAbiVersionException) What() string {
	return "ABI has an unsupported version"
}

func (e *UnsupportedAbiVersionException) AppendLog(l log.Message) {
	e.Elog = append(e.Elog, l)
}

func (e UnsupportedAbiVersionException) GetLog() log.Messages {
	return e.Elog
}

func (e UnsupportedAbiVersionException) TopMessage() string {
	for _, l := range e.Elog {
		if msg := l.GetMessage(); len(msg) > 0 {
			return msg
		}
	}
	return e.String()
}

func (e UnsupportedAbiVersionException) DetailMessage() string {
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

func (e UnsupportedAbiVersionException) String() string {
	return e.DetailMessage()
}

func (e UnsupportedAbiVersionException) MarshalJSON() ([]byte, error) {
	type Exception struct {
		Code int64  `json:"code"`
		Name string `json:"name"`
		What string `json:"what"`
	}

	except := Exception{
		Code: 3150016,
		Name: UnsupportedAbiVersionExceptionName,
		What: "ABI has an unsupported version",
	}

	return json.Marshal(except)
}

func (e UnsupportedAbiVersionException) Callback(f interface{}) bool {
	switch callback := f.(type) {
	case func(*UnsupportedAbiVersionException):
		callback(&e)
		return true
	case func(UnsupportedAbiVersionException):
		callback(e)
		return true
	default:
		return false
	}
}
