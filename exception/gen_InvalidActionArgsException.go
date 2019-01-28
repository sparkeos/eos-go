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

var InvalidActionArgsExceptionName = reflect.TypeOf(InvalidActionArgsException{}).Name()

type InvalidActionArgsException struct {
	_ActionValidateException
	Elog log.Messages
}

func NewInvalidActionArgsException(parent _ActionValidateException, message log.Message) *InvalidActionArgsException {
	return &InvalidActionArgsException{parent, log.Messages{message}}
}

func (e InvalidActionArgsException) Code() int64 {
	return 3050002
}

func (e InvalidActionArgsException) Name() string {
	return InvalidActionArgsExceptionName
}

func (e InvalidActionArgsException) What() string {
	return "Invalid Action Arguments"
}

func (e *InvalidActionArgsException) AppendLog(l log.Message) {
	e.Elog = append(e.Elog, l)
}

func (e InvalidActionArgsException) GetLog() log.Messages {
	return e.Elog
}

func (e InvalidActionArgsException) TopMessage() string {
	for _, l := range e.Elog {
		if msg := l.GetMessage(); len(msg) > 0 {
			return msg
		}
	}
	return e.String()
}

func (e InvalidActionArgsException) DetailMessage() string {
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

func (e InvalidActionArgsException) String() string {
	return e.DetailMessage()
}

func (e InvalidActionArgsException) MarshalJSON() ([]byte, error) {
	type Exception struct {
		Code int64  `json:"code"`
		Name string `json:"name"`
		What string `json:"what"`
	}

	except := Exception{
		Code: 3050002,
		Name: InvalidActionArgsExceptionName,
		What: "Invalid Action Arguments",
	}

	return json.Marshal(except)
}

func (e InvalidActionArgsException) Callback(f interface{}) bool {
	switch callback := f.(type) {
	case func(*InvalidActionArgsException):
		callback(&e)
		return true
	case func(InvalidActionArgsException):
		callback(e)
		return true
	default:
		return false
	}
}
