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

var ResolvedToMultiplePortsName = reflect.TypeOf(ResolvedToMultiplePorts{}).Name()

type ResolvedToMultiplePorts struct {
	_HttpException
	Elog log.Messages
}

func NewResolvedToMultiplePorts(parent _HttpException, message log.Message) *ResolvedToMultiplePorts {
	return &ResolvedToMultiplePorts{parent, log.Messages{message}}
}

func (e ResolvedToMultiplePorts) Code() int64 {
	return 3200003
}

func (e ResolvedToMultiplePorts) Name() string {
	return ResolvedToMultiplePortsName
}

func (e ResolvedToMultiplePorts) What() string {
	return "service resolved to multiple ports"
}

func (e *ResolvedToMultiplePorts) AppendLog(l log.Message) {
	e.Elog = append(e.Elog, l)
}

func (e ResolvedToMultiplePorts) GetLog() log.Messages {
	return e.Elog
}

func (e ResolvedToMultiplePorts) TopMessage() string {
	for _, l := range e.Elog {
		if msg := l.GetMessage(); len(msg) > 0 {
			return msg
		}
	}
	return e.String()
}

func (e ResolvedToMultiplePorts) DetailMessage() string {
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

func (e ResolvedToMultiplePorts) String() string {
	return e.DetailMessage()
}

func (e ResolvedToMultiplePorts) MarshalJSON() ([]byte, error) {
	type Exception struct {
		Code int64  `json:"code"`
		Name string `json:"name"`
		What string `json:"what"`
	}

	except := Exception{
		Code: 3200003,
		Name: ResolvedToMultiplePortsName,
		What: "service resolved to multiple ports",
	}

	return json.Marshal(except)
}

func (e ResolvedToMultiplePorts) Callback(f interface{}) bool {
	switch callback := f.(type) {
	case func(*ResolvedToMultiplePorts):
		callback(&e)
		return true
	case func(ResolvedToMultiplePorts):
		callback(e)
		return true
	default:
		return false
	}
}
