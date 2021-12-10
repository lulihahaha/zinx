package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"zinx/ziface"
)

// 定义全局参数，供其他模块使用

type GlobalObj struct {
	// server

	TcpServer ziface.IServer `json:"tcp_server"` // 当前Zinx全局的Server对象

	Host string `json:"host"` // 当前服务器主机监听的IP

	TcpPort int `json:"tcp_port"` // 当前服务器主机监听的端口号

	Name string `json:"name"` // 当前服务器的名称

	// zinx

	Version string `json:"version"` // 当前zinx的版本号

	MaxConn int `json:"max_conn"` // 当前服务器主机允许的最大链接数

	MaxPackageSize uint32 `json:"max_package_size"` // 当前zinx框架数据包的最大值
}

// 定义一个全局的对外的变量
var GlobalObject *GlobalObj

// 从zinx.json中去加载用于自定义的参数
func (g *GlobalObj) Reload() {
	data, err := os.Open("ZinxV0.4/zinx.json")
	if err != nil {
		fmt.Println("read file error:", err)
		panic(err)
	}
	defer data.Close()

	byteValue, err := ioutil.ReadAll(data)
	if err != nil {
		fmt.Println("read file error:", err)
		panic(err)
	}
	fmt.Println(string(byteValue))

	// 解析json数据到结构体中
	err = json.Unmarshal(byteValue, &GlobalObject)
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
