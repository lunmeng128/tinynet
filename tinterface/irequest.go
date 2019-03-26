package tinterface

type Request interface {
	GetConnection() Connection
	GetData() []byte
	GetMsgType() uint32
}
