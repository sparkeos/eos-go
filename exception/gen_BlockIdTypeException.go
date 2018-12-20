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

var BlockIdTypeExceptionName = reflect.TypeOf(BlockIdTypeException{}).Name()

type BlockIdTypeException struct {
	_ChainException
	Elog log.Messages
}

func NewBlockIdTypeException(parent _ChainException, message log.Message) *BlockIdTypeException {
	return &BlockIdTypeException{parent, log.Messages{message}}
}

func (e BlockIdTypeException) Code() int64 {
	return 3010008
}

func (e BlockIdTypeException) Name() string {
	return BlockIdTypeExceptionName
}

func (e BlockIdTypeException) What() string {
	return "Invalid block ID"
}

func (e *BlockIdTypeException) AppendLog(l log.Message) {
	e.Elog = append(e.Elog, l)
}

func (e BlockIdTypeException) GetLog() log.Messages {
	return e.Elog
}

func (e BlockIdTypeException) TopMessage() string {
	for _, l := range e.Elog {
		if msg := l.GetMessage(); msg != "" {
			return msg
		}
	}
	return e.String()
}

func (e BlockIdTypeException) DetailMessage() string {
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

func (e BlockIdTypeException) String() string {
	return e.DetailMessage()
}

func (e BlockIdTypeException) MarshalJSON() ([]byte, error) {
	type Exception struct {
		Code int64  `json:"code"`
		Name string `json:"name"`
		What string `json:"what"`
	}

	except := Exception{
		Code: 3010008,
		Name: BlockIdTypeExceptionName,
		What: "Invalid block ID",
	}

	return json.Marshal(except)
}

func (e BlockIdTypeException) Callback(f interface{}) bool {
	switch callback := f.(type) {
	case func(*BlockIdTypeException):
		callback(&e)
		return true
	case func(BlockIdTypeException):
		callback(e)
		return true
	default:
		return false
	}
}
