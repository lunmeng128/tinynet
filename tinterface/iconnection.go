package tinterface

import "net"

type Connection interface {
	Start()
	Stop()
	GetTCPConnection() *net.TCPConn
	GetConnID() uint32
	GetClientAddr() net.Addr
	SendMsg(msgType uint32, data []byte) error
	SetProperty(key string, value interface{})
	GetProperty(key string) (interface{}, error)
	DelProperty(key string)
}
