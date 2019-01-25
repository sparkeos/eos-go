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

var ActionDataAndStructMismatchName = reflect.TypeOf(ActionDataAndStructMismatch{}).Name()

type ActionDataAndStructMismatch struct {
	_ActionValidateException
	Elog log.Messages
}

func NewActionDataAndStructMismatch(parent _ActionValidateException, message log.Message) *ActionDataAndStructMismatch {
	return &ActionDataAndStructMismatch{parent, log.Messages{message}}
}

func (e ActionDataAndStructMismatch) Code() int64 {
	return 3050006
}

func (e ActionDataAndStructMismatch) Name() string {
	return ActionDataAndStructMismatchName
}

func (e ActionDataAndStructMismatch) What() string {
	return "Mismatch between action data and its struct"
}

func (e *ActionDataAndStructMismatch) AppendLog(l log.Message) {
	e.Elog = append(e.Elog, l)
}

func (e ActionDataAndStructMismatch) GetLog() log.Messages {
	return e.Elog
}

func (e ActionDataAndStructMismatch) TopMessage() string {
	for _, l := range e.Elog {
		if msg := l.GetMessage(); len(msg) > 0 {
			return msg
		}
	}
	return e.String()
}

func (e ActionDataAndStructMismatch) DetailMessage() string {
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

func (e ActionDataAndStructMismatch) String() string {
	return e.DetailMessage()
}

func (e ActionDataAndStructMismatch) MarshalJSON() ([]byte, error) {
	type Exception struct {
		Code int64  `json:"code"`
		Name string `json:"name"`
		What string `json:"what"`
	}

	except := Exception{
		Code: 3050006,
		Name: ActionDataAndStructMismatchName,
		What: "Mismatch between action data and its struct",
	}

	return json.Marshal(except)
}

func (e ActionDataAndStructMismatch) Callback(f interface{}) bool {
	switch callback := f.(type) {
	case func(*ActionDataAndStructMismatch):
		callback(&e)
		return true
	case func(ActionDataAndStructMismatch):
		callback(e)
		return true
	default:
		return false
	}
}
