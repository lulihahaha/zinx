package ziface

// 把客户端请求的链接信息和请求的数据封装到一个request中

type IRequest interface {
	// 得到当前链接
	GetConnection() IConnection

	// 得到请求的消息数据
	GetData() []byte
}
