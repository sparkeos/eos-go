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

var DbApiExceptionName = reflect.TypeOf(DbApiException{}).Name()

type DbApiException struct {
	_ContractApiException
	Elog log.Messages
}

func NewDbApiException(parent _ContractApiException, message log.Message) *DbApiException {
	return &DbApiException{parent, log.Messages{message}}
}

func (e DbApiException) Code() int64 {
	return 3230002
}

func (e DbApiException) Name() string {
	return DbApiExceptionName
}

func (e DbApiException) What() string {
	return "Database API exception"
}

func (e *DbApiException) AppendLog(l log.Message) {
	e.Elog = append(e.Elog, l)
}

func (e DbApiException) GetLog() log.Messages {
	return e.Elog
}

func (e DbApiException) TopMessage() string {
	for _, l := range e.Elog {
		if msg := l.GetMessage(); msg != "" {
			return msg
		}
	}
	return e.String()
}

func (e DbApiException) DetailMessage() string {
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

func (e DbApiException) String() string {
	return e.DetailMessage()
}

func (e DbApiException) MarshalJSON() ([]byte, error) {
	type Exception struct {
		Code int64  `json:"code"`
		Name string `json:"name"`
		What string `json:"what"`
	}

	except := Exception{
		Code: 3230002,
		Name: DbApiExceptionName,
		What: "Database API exception",
	}

	return json.Marshal(except)
}

func (e DbApiException) Callback(f interface{}) bool {
	switch callback := f.(type) {
	case func(*DbApiException):
		callback(&e)
		return true
	case func(DbApiException):
		callback(e)
		return true
	default:
		return false
	}
}