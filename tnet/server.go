package tnet

import (
	"fmt"
	"net"
	"tinynet/tinterface"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
	ConnMgr   tinterface.ConnManager
	MsgHandle tinterface.MsgHandle
}

func NewServer() tinterface.Server {
	ConfigObj.LoadConfig()
	s := &Server{
		Name:      ConfigObj.Name,
		IPVersion: "tcp4",
		IP:        ConfigObj.Host,
		Port:      ConfigObj.Port,
		ConnMgr:   NewConnManager(),
		MsgHandle: NewMsgHandle(),
	}
	return s
}

func (s *Server) Start() {
	fmt.Printf("[START] Server name: %s,listenner at IP: %s, Port %d is starting\n", s.Name, s.IP, s.Port)
	fmt.Printf("[tinynet] Version: %s, MaxConn: %d, MaxPacketSize: %d\n",
		ConfigObj.Version,
		ConfigObj.MaxConn,
		ConfigObj.MaxPacketSize)

	go func() {
		address := fmt.Sprintf("%s:%d", s.IP, s.Port)
		tcpAddr, err := net.ResolveTCPAddr(s.IPVersion, address)
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}
		listenner, err := net.ListenTCP(s.IPVersion, tcpAddr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}

		fmt.Println("start tinynet server  ", s.Name, address, " succ, now listenning...")

		//session ID 后期规划一个ID下发
		var cid uint32
		cid = 0

		for {
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println(" Accept Tcp error ", err.Error())
				continue
			}

			fmt.Println("Get conn remote addr = ", conn.RemoteAddr().String())

			//3.2 设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接
			if s.ConnMgr.Len() >= ConfigObj.MaxConn {
				conn.Close()
				continue
			}

			//3.3 处理该新连接请求的 业务 方法， 此时应该有 handler 和 conn是绑定的
			dealConn := NewConntion(s, conn, cid, s.MsgHandle)
			cid++
			//3.4 启动当前链接的处理业务
			go dealConn.Start()
		}
	}()
	select {}
}

func (s *Server) Stop() {
	s.GetConnManger().ClearConn()
}

func (s *Server) Serve() {
	s.Start()
}

func (s *Server) GetConnManger() tinterface.ConnManager {
	return s.ConnMgr
}

func (s *Server) AddRouter(msgType uint32, router tinterface.Router) {
	s.MsgHandle.AddRouter(msgType, router)
}
