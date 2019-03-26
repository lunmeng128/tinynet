package tnet

import (
	"tinynet/tinterface"
)

type Request struct {
	conn tinterface.Connection
	msg  tinterface.Message
}

func (r *Request) GetConnection() tinterface.Connection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

func (r *Request) GetMsgType() uint32 {
	return r.msg.GetMsgType()
}
