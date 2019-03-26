package tinterface

type Message interface {
	GetDataLen() uint32
	GetMsgType() uint32
	GetData() []byte

	SetMsgType(uint32)
	SetData([]byte)
	SetDataLen(uint32)
}
