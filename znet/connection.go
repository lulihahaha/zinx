package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

// 链接模块
type Connection struct {
	// 当前链接的Socket TCP套接字
	Conn *net.TCPConn

	// 链接的ID
	ConnID uint32

	// 当前的链接状态
	IsClosed bool

	// 当前链接所绑定的处理业务方法API
	HandleAPI ziface.HandleFunc

	// 告知当前链接已经退出/停止的channel
	ExistChan chan bool
}

// 链接的读业务方法
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Println("connID = ", c.ConnID, "Reader is exit,remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {
		// 读取客户端的数据到buf中
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("receive buf error", err)
			continue
		}

		// 调用当前链接所绑定的Handle API
		if err := c.HandleAPI(c.Conn, buf, cnt); err != nil {
			fmt.Println("ConnID = ", c.ConnID, "handle is error", err)
			break
		}
	}
}

func (c *Connection) Start() {
	fmt.Println("Connection Start,ConnID = ", c.ConnID)

	// 启动从当前链接的读数据的业务
	go c.StartReader()

	// TODO 启动从当前连接写数据的业务
}

func (c *Connection) Stop() {
	fmt.Println("Connection Stop,ConnID = ", c.ConnID)

	// 如果当前链接已经关闭
	if c.IsClosed == true {
		return
	}
	c.IsClosed = true

	// 关闭socket链接
	c.Conn.Close()

	// 回收资源
	close(c.ExistChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(data []byte) error {
	return nil
}

// 初始化链接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, callback_api ziface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		HandleAPI: callback_api,
		IsClosed:  false,
		ExistChan: make(chan bool, 1),
	}

	return c
}
