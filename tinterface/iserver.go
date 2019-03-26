package tinterface

type Server interface {
	Start()
	Stop()
	Serve()
	GetConnManger() ConnManager
	AddRouter(msgType uint32, router Router)
}
