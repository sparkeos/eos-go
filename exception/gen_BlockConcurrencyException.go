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

var BlockConcurrencyExceptionName = reflect.TypeOf(BlockConcurrencyException{}).Name()

type BlockConcurrencyException struct {
	_BlockValidateException
	Elog log.Messages
}

func NewBlockConcurrencyException(parent _BlockValidateException, message log.Message) *BlockConcurrencyException {
	return &BlockConcurrencyException{parent, log.Messages{message}}
}

func (e BlockConcurrencyException) Code() int64 {
	return 3030003
}

func (e BlockConcurrencyException) Name() string {
	return BlockConcurrencyExceptionName
}

func (e BlockConcurrencyException) What() string {
	return "Block does not guarantee concurrent execution without conflicts"
}

func (e *BlockConcurrencyException) AppendLog(l log.Message) {
	e.Elog = append(e.Elog, l)
}

func (e BlockConcurrencyException) GetLog() log.Messages {
	return e.Elog
}

func (e BlockConcurrencyException) TopMessage() string {
	for _, l := range e.Elog {
		if msg := l.GetMessage(); msg != "" {
			return msg
		}
	}
	return e.String()
}

func (e BlockConcurrencyException) DetailMessage() string {
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

func (e BlockConcurrencyException) String() string {
	return e.DetailMessage()
}

func (e BlockConcurrencyException) MarshalJSON() ([]byte, error) {
	type Exception struct {
		Code int64  `json:"code"`
		Name string `json:"name"`
		What string `json:"what"`
	}

	except := Exception{
		Code: 3030003,
		Name: BlockConcurrencyExceptionName,
		What: "Block does not guarantee concurrent execution without conflicts",
	}

	return json.Marshal(except)
}

func (e BlockConcurrencyException) Callback(f interface{}) bool {
	switch callback := f.(type) {
	case func(*BlockConcurrencyException):
		callback(&e)
		return true
	case func(BlockConcurrencyException):
		callback(e)
		return true
	default:
		return false
	}
}
