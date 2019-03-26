package main

import (
	"tinynet/tnet"
	"tinynet/demo/server/router"
)

func main() {
	//1 创建一个server句柄
	tnet.Init()
	s := tnet.NewServer()
	//3 开启服务
	s.AddRouter(0,&router.Hello{})
	s.Serve()
}
