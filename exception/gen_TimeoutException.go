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

var TimeoutExceptionName = reflect.TypeOf(TimeoutException{}).Name()

type TimeoutException struct {
	Exception
	Elog log.Messages
}

func NewTimeoutException(parent Exception, message log.Message) *TimeoutException {
	return &TimeoutException{parent, log.Messages{message}}
}

func (e TimeoutException) Code() int64 {
	return TimeoutExceptionCode
}

func (e TimeoutException) Name() string {
	return TimeoutExceptionName
}

func (e TimeoutException) What() string {
	return "Timeout"
}

func (e *TimeoutException) AppendLog(l log.Message) {
	e.Elog = append(e.Elog, l)
}

func (e TimeoutException) GetLog() log.Messages {
	return e.Elog
}

func (e TimeoutException) TopMessage() string {
	for _, l := range e.Elog {
		if msg := l.GetMessage(); len(msg) > 0 {
			return msg
		}
	}
	return e.String()
}

func (e TimeoutException) DetailMessage() string {
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

func (e TimeoutException) String() string {
	return e.DetailMessage()
}

func (e TimeoutException) MarshalJSON() ([]byte, error) {
	type Exception struct {
		Code int64  `json:"code"`
		Name string `json:"name"`
		What string `json:"what"`
	}

	except := Exception{
		Code: TimeoutExceptionCode,
		Name: TimeoutExceptionName,
		What: "Timeout",
	}

	return json.Marshal(except)
}

func (e TimeoutException) Callback(f interface{}) bool {
	switch callback := f.(type) {
	case func(*TimeoutException):
		callback(&e)
		return true
	case func(TimeoutException):
		callback(e)
		return true
	default:
		return false
	}
}
