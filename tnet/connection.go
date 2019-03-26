package tnet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"tinynet/tinterface"
)

type Connection struct {
	TcpServer tinterface.Server
	TcpConn   *net.TCPConn
	//当前连接的ID 也可以称作为SessionID，ID全局唯一
	ConnID uint32
	//当前连接的关闭状态
	isClosed bool
	//消息管理MsgType和对应处理方法的消息管理模块
	MsgHandler tinterface.MsgHandle
	//告知该链接已经退出/停止的channel
	ExitBuffChan chan bool
	//无缓冲管道，用于读、写两个goroutine之间的消息通信
	msgChan chan []byte
	//有关冲管道，用于读、写两个goroutine之间的消息通信
	msgBuffChan chan []byte

	//链接属性
	property map[string]interface{}
	//保护链接属性修改的锁
	propertyLock sync.RWMutex
}

func NewConntion(server tinterface.Server, conn *net.TCPConn, connID uint32, msgHandle tinterface.MsgHandle) Connection {
	return Connection{
		TcpServer:    server,
		TcpConn:      conn,
		ConnID:       connID,
		MsgHandler:   msgHandle,
		isClosed:     false,
		ExitBuffChan: make(chan bool, 1),
		msgChan:      make(chan []byte),
		msgBuffChan:  make(chan []byte, ConfigObj.MaxMsgChanLen),
		property:     make(map[string]interface{}),
	}
}

func (c *Connection) Start() {
	go c.StartReader()
	go c.StartWriter()
}

func (c *Connection) StartWriter() {
	fmt.Println("[Writer Goroutine is running]")
	defer func() {
		fmt.Println(c.GetClientAddr().String(), "[conn Writer exit!]")
		c.Stop()
	}()
	for {
		select {
		case msg := <-c.msgChan:
			if _, err := c.TcpConn.Write(msg); err != nil {
				fmt.Println(" msgChan write data error", err.Error())
				return
			}
		case msg, ok := <-c.msgBuffChan:
			if ok {
				if _, err := c.TcpConn.Write(msg); err != nil {
					fmt.Println(" msgBuffChan write data error", err.Error())
					return
				}
			} else {
				fmt.Println("msgBuffChan is Closed")
				break
			}
		case <-c.ExitBuffChan:
			return
		}
	}
}

func (c *Connection) StartReader() {
	fmt.Println("[Reader Goroutine is running]")
	defer func() {
		fmt.Println(c.GetClientAddr().String(), "[conn Reader exit!]")
		c.Stop()
	}()
	for {
		dataCodec := NewCodec()
		HeadData := make([]byte, dataCodec.GetHeadLen())
		if _, err := io.ReadFull(c.GetTcpConnection(), HeadData); err != nil {
			//error 返回eof 则退出当前连接
			break
		}

		msg, err := dataCodec.UnPack(HeadData)
		if err != nil {
			fmt.Println("unpack error", err)
			break
		}

		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(c.GetTcpConnection(), data); err != nil {
				fmt.Println("read msg data error ", err)
				break
			}
		}

		msg.SetData(data)

		request := &Request{
			conn: c,
			msg:  msg,
		}
		//根据不同的消息ID 处理不同的消息类型 类型消TYPE 类型
		go c.MsgHandler.DoMsgHandler(request)
	}
}

func (c *Connection) GetTcpConnection() *net.TCPConn {
	return c.TcpConn
}

func (c *Connection) SendMsg(msgType uint32, data []byte) error {
	if c.isClosed {
		return errors.New("Connection closed when send msg")
	}
	dataCodec := NewCodec()
	bt, err := dataCodec.Pack(NewMessage(msgType, data))
	if err != nil {
		return errors.New("codec data error")
	}
	c.msgChan <- bt
	return nil
}
func (c *Connection) SetProperty(key string, value interface{}) {

}
func (c *Connection) GetProperty(key string) (interface{}, error) {
	return nil, nil
}

func (c *Connection) DelProperty(key string) {

}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) GetClientAddr() net.Addr {
	return c.TcpConn.RemoteAddr()
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.TcpConn
}

func (c *Connection) Stop() {
	if c.isClosed {
		return
	}
	c.isClosed = true
	c.TcpConn.Close()
	c.TcpServer.GetConnManger().Del(c)
	c.ExitBuffChan <- true

	close(c.ExitBuffChan)
	close(c.msgBuffChan)
}
