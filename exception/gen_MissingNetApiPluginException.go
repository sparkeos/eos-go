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

var MissingNetApiPluginExceptionName = reflect.TypeOf(MissingNetApiPluginException{}).Name()

type MissingNetApiPluginException struct {
	_PluginException
	Elog log.Messages
}

func NewMissingNetApiPluginException(parent _PluginException, message log.Message) *MissingNetApiPluginException {
	return &MissingNetApiPluginException{parent, log.Messages{message}}
}

func (e MissingNetApiPluginException) Code() int64 {
	return 3110004
}

func (e MissingNetApiPluginException) Name() string {
	return MissingNetApiPluginExceptionName
}

func (e MissingNetApiPluginException) What() string {
	return "Missing Net API Plugin"
}

func (e *MissingNetApiPluginException) AppendLog(l log.Message) {
	e.Elog = append(e.Elog, l)
}

func (e MissingNetApiPluginException) GetLog() log.Messages {
	return e.Elog
}

func (e MissingNetApiPluginException) TopMessage() string {
	for _, l := range e.Elog {
		if msg := l.GetMessage(); msg != "" {
			return msg
		}
	}
	return e.String()
}

func (e MissingNetApiPluginException) DetailMessage() string {
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

func (e MissingNetApiPluginException) String() string {
	return e.DetailMessage()
}

func (e MissingNetApiPluginException) MarshalJSON() ([]byte, error) {
	type Exception struct {
		Code int64  `json:"code"`
		Name string `json:"name"`
		What string `json:"what"`
	}

	except := Exception{
		Code: 3110004,
		Name: MissingNetApiPluginExceptionName,
		What: "Missing Net API Plugin",
	}

	return json.Marshal(except)
}

func (e MissingNetApiPluginException) Callback(f interface{}) bool {
	switch callback := f.(type) {
	case func(*MissingNetApiPluginException):
		callback(&e)
		return true
	case func(MissingNetApiPluginException):
		callback(e)
		return true
	default:
		return false
	}
}
