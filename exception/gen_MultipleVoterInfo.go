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

var MultipleVoterInfoName = reflect.TypeOf(MultipleVoterInfo{}).Name()

type MultipleVoterInfo struct {
	_MiscException
	Elog log.Messages
}

func NewMultipleVoterInfo(parent _MiscException, message log.Message) *MultipleVoterInfo {
	return &MultipleVoterInfo{parent, log.Messages{message}}
}

func (e MultipleVoterInfo) Code() int64 {
	return 3100007
}

func (e MultipleVoterInfo) Name() string {
	return MultipleVoterInfoName
}

func (e MultipleVoterInfo) What() string {
	return "Multiple voter info detected"
}

func (e *MultipleVoterInfo) AppendLog(l log.Message) {
	e.Elog = append(e.Elog, l)
}

func (e MultipleVoterInfo) GetLog() log.Messages {
	return e.Elog
}

func (e MultipleVoterInfo) TopMessage() string {
	for _, l := range e.Elog {
		if msg := l.GetMessage(); len(msg) > 0 {
			return msg
		}
	}
	return e.String()
}

func (e MultipleVoterInfo) DetailMessage() string {
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

func (e MultipleVoterInfo) String() string {
	return e.DetailMessage()
}

func (e MultipleVoterInfo) MarshalJSON() ([]byte, error) {
	type Exception struct {
		Code int64  `json:"code"`
		Name string `json:"name"`
		What string `json:"what"`
	}

	except := Exception{
		Code: 3100007,
		Name: MultipleVoterInfoName,
		What: "Multiple voter info detected",
	}

	return json.Marshal(except)
}

func (e MultipleVoterInfo) Callback(f interface{}) bool {
	switch callback := f.(type) {
	case func(*MultipleVoterInfo):
		callback(&e)
		return true
	case func(MultipleVoterInfo):
		callback(e)
		return true
	default:
		return false
	}
}
