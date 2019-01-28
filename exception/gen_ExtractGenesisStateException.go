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

var ExtractGenesisStateExceptionName = reflect.TypeOf(ExtractGenesisStateException{}).Name()

type ExtractGenesisStateException struct {
	_MiscException
	Elog log.Messages
}

func NewExtractGenesisStateException(parent _MiscException, message log.Message) *ExtractGenesisStateException {
	return &ExtractGenesisStateException{parent, log.Messages{message}}
}

func (e ExtractGenesisStateException) Code() int64 {
	return 3100005
}

func (e ExtractGenesisStateException) Name() string {
	return ExtractGenesisStateExceptionName
}

func (e ExtractGenesisStateException) What() string {
	return "Extracted genesis state from blocks.log"
}

func (e *ExtractGenesisStateException) AppendLog(l log.Message) {
	e.Elog = append(e.Elog, l)
}

func (e ExtractGenesisStateException) GetLog() log.Messages {
	return e.Elog
}

func (e ExtractGenesisStateException) TopMessage() string {
	for _, l := range e.Elog {
		if msg := l.GetMessage(); len(msg) > 0 {
			return msg
		}
	}
	return e.String()
}

func (e ExtractGenesisStateException) DetailMessage() string {
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

func (e ExtractGenesisStateException) String() string {
	return e.DetailMessage()
}

func (e ExtractGenesisStateException) MarshalJSON() ([]byte, error) {
	type Exception struct {
		Code int64  `json:"code"`
		Name string `json:"name"`
		What string `json:"what"`
	}

	except := Exception{
		Code: 3100005,
		Name: ExtractGenesisStateExceptionName,
		What: "Extracted genesis state from blocks.log",
	}

	return json.Marshal(except)
}

func (e ExtractGenesisStateException) Callback(f interface{}) bool {
	switch callback := f.(type) {
	case func(*ExtractGenesisStateException):
		callback(&e)
		return true
	case func(ExtractGenesisStateException):
		callback(e)
		return true
	default:
		return false
	}
}
