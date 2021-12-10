package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

// 基于Zinx框架来开发的服务端应用程序

func main() {
	// 创建一个Server句柄，使用Zinx的api
	s := znet.NewServer()
	// 给zinx框架添加router
	s.AddRouter(&PingRouter{})
	// 启动Server
	s.Serve()
}

// ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

func (this *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("call router prehandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Println("call back before ping error")
	}
}

func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("call router handle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping.. ping...\n"))
	if err != nil {
		fmt.Println("call back ping error")
	}
}

func (this *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("call router posthandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping...\n"))
	if err != nil {
		fmt.Println("call back after ping error")
	}
}
