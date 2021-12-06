package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("client start...")
	time.Sleep(1 * time.Second)

	// 1.直接连接远程服务器，得到一个conn连接
	conn, err := net.Dial("tcp", "139.198.190.28:8999")
	if err != nil {
		fmt.Println("client start error,exit! ")
		return
	}

	for {
		// 2.调用write写数据
		_, err := conn.Write([]byte("Hello Zinx v0.1..."))
		if err != nil {
			fmt.Println("conn write error:", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error", err)
			return
		}

		fmt.Printf("server call back:%s,cnt=%d\n", buf, cnt)

		time.Sleep(1 * time.Second)
	}
}
