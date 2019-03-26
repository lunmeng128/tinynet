package tnet

type Message struct {
	DataLen uint32 //消息的长度
	MsgType uint32 //消息的ID
	Data    []byte //消息的内容
}

func NewMessage(msgType uint32, data []byte) *Message {
	return &Message{
		DataLen: uint32(len(data)),
		MsgType: msgType,
		Data:    data,
	}
}

func (m *Message) GetDataLen() uint32 {
	return m.DataLen
}
func (m *Message) GetMsgType() uint32 {
	return m.MsgType
}
func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetMsgType(msgType uint32) {
	m.MsgType = msgType
}

func (m *Message) SetData(data []byte) {
	m.Data = data
}

func (m *Message) SetDataLen(dataLen uint32) {
	m.DataLen = dataLen
}
