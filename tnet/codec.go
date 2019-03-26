package tnet

import (
	"bytes"
	"encoding/binary"
	"tinynet/tinterface"
)

type Codec struct {
}

// 协议由包头+包体组成
// 包头 消息id (4字节) + 包体的长度(4字节)
// 包体 固定长度的byte字节
func NewCodec() *Codec {
	return &Codec{}
}

func (c *Codec) GetHeadLen() int {
	return 8
}

// Pack 打包
func (c *Codec) Pack(msg tinterface.Message) ([]byte, error) {
	dataBuff := bytes.NewBuffer([]byte{})
	//写入data包的长度
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}
	//写入msg type
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgType()); err != nil {
		return nil, err
	}
	//写入消息体 protobuf
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return dataBuff.Bytes(), nil
}

// UnPack 数据解码
func (c *Codec) UnPack(binaryData []byte) (tinterface.Message, error) {
	dataBuff := bytes.NewReader(binaryData)
	msg := &Message{}
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.MsgType); err != nil {
		return nil, err
	}
	return msg, nil
}
