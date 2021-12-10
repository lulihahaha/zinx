package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"zinx/ziface"
)

// 定义全局参数，供其他模块使用

type GlobalObj struct {
	// server

	TcpServer ziface.IServer // 当前Zinx全局的Server对象

	Host string // 当前服务器主机监听的IP

	TcpPort int // 当前服务器主机监听的端口号

	Name string // 当前服务器的名称

	// zinx

	Version string // 当前zinx的版本号

	MaxConn int // 当前服务器主机允许的最大链接数

	MaxPackageSize uint32 // 当前zinx框架数据包的最大值
}

// 定义一个全局的对外的变量
var GlobalObject *GlobalObj

// 从zinx.json中去加载用于自定义的参数
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		fmt.Println("read file error:", err)
		panic(err)
	}

	// 解析json数据到结构体中
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

// 提供一个init方法，初始化当前的GlobalObject
func init() {
	// 如果配置文件没有加载，就使用默认值
	GlobalObject = &GlobalObj{
		Name:           "ZinxServerApp",
		Version:        "v0.4",
		TcpPort:        8999,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}

	GlobalObject.Reload()
}
