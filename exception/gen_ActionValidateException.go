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

var ActionValidateExceptionName = reflect.TypeOf(ActionValidateException{}).Name()

type ActionValidateException struct {
	_ActionValidateException
	Elog log.Messages
}

func NewActionValidateException(parent _ActionValidateException, message log.Message) *ActionValidateException {
	return &ActionValidateException{parent, log.Messages{message}}
}

func (e ActionValidateException) Code() int64 {
	return 3050000
}

func (e ActionValidateException) Name() string {
	return ActionValidateExceptionName
}

func (e ActionValidateException) What() string {
	return "Transaction exceeded the current CPU usage limit imposed on the transaction"
}

func (e *ActionValidateException) AppendLog(l log.Message) {
	e.Elog = append(e.Elog, l)
}

func (e ActionValidateException) GetLog() log.Messages {
	return e.Elog
}

func (e ActionValidateException) TopMessage() string {
	for _, l := range e.Elog {
		if msg := l.GetMessage(); len(msg) > 0 {
			return msg
		}
	}
	return e.String()
}

func (e ActionValidateException) DetailMessage() string {
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

func (e ActionValidateException) String() string {
	return e.DetailMessage()
}

func (e ActionValidateException) MarshalJSON() ([]byte, error) {
	type Exception struct {
		Code int64  `json:"code"`
		Name string `json:"name"`
		What string `json:"what"`
	}

	except := Exception{
		Code: 3050000,
		Name: ActionValidateExceptionName,
		What: "Transaction exceeded the current CPU usage limit imposed on the transaction",
	}

	return json.Marshal(except)
}

func (e ActionValidateException) Callback(f interface{}) bool {
	switch callback := f.(type) {
	case func(*ActionValidateException):
		callback(&e)
		return true
	case func(ActionValidateException):
		callback(e)
		return true
	default:
		return false
	}
}
