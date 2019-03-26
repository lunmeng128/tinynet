package tnet

import (
	"fmt"
	"strconv"
	"tinynet/tinterface"
)

type MsgHandle struct {
	Apis map[uint32]tinterface.Router
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]tinterface.Router, 0),
	}
}

func (m *MsgHandle) StartWorkerPool() {

}

func (m *MsgHandle) DoMsgHandler(request tinterface.Request) {
	handle, ok := m.Apis[request.GetMsgType()]
	if !ok {
		fmt.Println(" api not found", request.GetMsgType())
		return
	}
	handle.Handle(request)

}

func (m *MsgHandle) AddRouter(msgType uint32, router tinterface.Router) {
	if _, ok := m.Apis[msgType]; ok {
		panic("repeated api , msgType = " + strconv.Itoa(int(msgType)))
	}
	//2 添加msg与api的绑定关系
	m.Apis[msgType] = router
	fmt.Println("Add api msgType = ", msgType)
}
