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

var TxCpuUsageExceededName = reflect.TypeOf(TxCpuUsageExceeded{}).Name()

type TxCpuUsageExceeded struct {
	_ResourceExhaustedException
	Elog log.Messages
}

func NewTxCpuUsageExceeded(parent _ResourceExhaustedException, message log.Message) *TxCpuUsageExceeded {
	return &TxCpuUsageExceeded{parent, log.Messages{message}}
}

func (e TxCpuUsageExceeded) Code() int64 {
	return 3080004
}

func (e TxCpuUsageExceeded) Name() string {
	return TxCpuUsageExceededName
}

func (e TxCpuUsageExceeded) What() string {
	return "Transaction exceeded the current CPU usage limit imposed on the transaction"
}

func (e *TxCpuUsageExceeded) AppendLog(l log.Message) {
	e.Elog = append(e.Elog, l)
}

func (e TxCpuUsageExceeded) GetLog() log.Messages {
	return e.Elog
}

func (e TxCpuUsageExceeded) TopMessage() string {
	for _, l := range e.Elog {
		if msg := l.GetMessage(); len(msg) > 0 {
			return msg
		}
	}
	return e.String()
}

func (e TxCpuUsageExceeded) DetailMessage() string {
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

func (e TxCpuUsageExceeded) String() string {
	return e.DetailMessage()
}

func (e TxCpuUsageExceeded) MarshalJSON() ([]byte, error) {
	type Exception struct {
		Code int64  `json:"code"`
		Name string `json:"name"`
		What string `json:"what"`
	}

	except := Exception{
		Code: 3080004,
		Name: TxCpuUsageExceededName,
		What: "Transaction exceeded the current CPU usage limit imposed on the transaction",
	}

	return json.Marshal(except)
}

func (e TxCpuUsageExceeded) Callback(f interface{}) bool {
	switch callback := f.(type) {
	case func(*TxCpuUsageExceeded):
		callback(&e)
		return true
	case func(TxCpuUsageExceeded):
		callback(e)
		return true
	default:
		return false
	}
}
