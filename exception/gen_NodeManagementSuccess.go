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

var NodeManagementSuccessName = reflect.TypeOf(NodeManagementSuccess{}).Name()

type NodeManagementSuccess struct {
	_MiscException
	Elog log.Messages
}

func NewNodeManagementSuccess(parent _MiscException, message log.Message) *NodeManagementSuccess {
	return &NodeManagementSuccess{parent, log.Messages{message}}
}

func (e NodeManagementSuccess) Code() int64 {
	return 3100009
}

func (e NodeManagementSuccess) Name() string {
	return NodeManagementSuccessName
}

func (e NodeManagementSuccess) What() string {
	return "Node management operation successfully executed"
}

func (e *NodeManagementSuccess) AppendLog(l log.Message) {
	e.Elog = append(e.Elog, l)
}

func (e NodeManagementSuccess) GetLog() log.Messages {
	return e.Elog
}

func (e NodeManagementSuccess) TopMessage() string {
	for _, l := range e.Elog {
		if msg := l.GetMessage(); msg != "" {
			return msg
		}
	}
	return e.String()
}

func (e NodeManagementSuccess) DetailMessage() string {
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

func (e NodeManagementSuccess) String() string {
	return e.DetailMessage()
}

func (e NodeManagementSuccess) MarshalJSON() ([]byte, error) {
	type Exception struct {
		Code int64  `json:"code"`
		Name string `json:"name"`
		What string `json:"what"`
	}

	except := Exception{
		Code: 3100009,
		Name: NodeManagementSuccessName,
		What: "Node management operation successfully executed",
	}

	return json.Marshal(except)
}

func (e NodeManagementSuccess) Callback(f interface{}) bool {
	switch callback := f.(type) {
	case func(*NodeManagementSuccess):
		callback(&e)
		return true
	case func(NodeManagementSuccess):
		callback(e)
		return true
	default:
		return false
	}
}
