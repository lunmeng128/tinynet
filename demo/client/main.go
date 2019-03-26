package main

import (
	"fmt"
	"net"
	"time"
	"tinynet/tnet"
	"io"
)

/*
	模拟客户端
*/
func main() {

	fmt.Println("Client Test ... start")
	//3秒之后发起测试请求，给服务端开启服务的机会
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "0.0.0.0:9999")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	for n := 3; n >= 0; n-- {
		dp := tnet.NewCodec()
		msg, _ := dp.Pack(tnet.NewMessage(0, []byte("tinynet Client Test Message")))
		_, err := conn.Write(msg)
		if err != nil {
			fmt.Println(err)
		}
		//先读出流中的head部分
		headData := make([]byte, dp.GetHeadLen())
		_, err = io.ReadFull(conn, headData) //ReadFull 会把msg填充满为止
		if err != nil {
			fmt.Println("read head error")
			break
		}
		//将headData字节流 拆包到msg中
		msgHead, err := dp.UnPack(headData)
		if err != nil {
			fmt.Println("server unpack err:", err)
			return
		}

		if msgHead.GetDataLen() > 0 {
			//msg 是有data数据的，需要再次读取data数据
			msg := msgHead.(*tnet.Message)
			msg.Data = make([]byte, msg.GetDataLen())

			//根据dataLen从io中读取字节流
			_, err := io.ReadFull(conn, msg.Data)
			if err != nil {
				fmt.Println("server unpack data err:", err)
				return
			}
			fmt.Println("==> Recv Msg: Type=", msg.MsgType, ", len=", msg.DataLen, ", data=", string(msg.Data))
		}

		time.Sleep(3 * time.Second)
	}
	select {}
	conn.Close()
}
