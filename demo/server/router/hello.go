package router

import (
	"tinynet/tinterface"
	"tinynet/tnet"
)

type Hello struct {
	tnet.BaseRouter
}

func (h *Hello) Handle(request tinterface.Request) {
	request.GetConnection().SendMsg(request.GetMsgType(),[]byte("hello"))
}
