package znet

import (
	"fmt"
	"net"
	"zinx/utils"
	"zinx/ziface"
)

// iServer的接口实现，定义一个Server的服务器模块
type Server struct {
	// 服务器的名称
	Name string
	// 服务器绑定的ip版本
	IPVersion string
	// 服务器监听的IP
	IP string
	// 服务器监听的端口
	Port int
	// 当前的server添加一个router,暂时一个服务器只能添加一个router
	Router ziface.IRouter
}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("Add router successfully!")
}

func (s *Server) Start() {
	fmt.Printf("[Start] Server Listenner at IP:%s,Port %d, is starting\n", s.IP, s.Port)
	fmt.Printf("[Zinx] SeverName is %s, SeverVersion is %s\n", utils.GlobalObject.Name, utils.GlobalObject.Version)

	go func() {
		// 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error:", err)
			return
		}

		// 监听服务器的地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen ", s.IPVersion, "err ", err)
		}

		fmt.Println("start Zinx server successfully ", s.Name, " successfully, Listenning...")
		var cid uint32
		cid = 0

		// 阻塞等待客户端连接，处理客户端的链接业务（读写
		for {
			// 如果有客户端连接过来，阻塞会返回
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}

			// 将处理新连接的业务方法和conn进心绑定，得到我们的连接模块
			dealConn := NewConnection(conn, cid, s.Router)
			cid++

			// 启动当前的链接业务
			go dealConn.Start()

		}

	}()
}

func (s *Server) Stop() {
	// TODO
}

func (s *Server) Serve() {
	// 启动server的服务功能
	s.Start()

	// TODO

	// 阻塞状态
	select {}
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.TcpPort,
		Router:    nil,
	}
	return s
}
