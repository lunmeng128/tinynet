package tinterface

type MsgHandle interface {
	StartWorkerPool()
	DoMsgHandler(request Request)
	AddRouter(msgTypemsgType uint32, router Router)
}
