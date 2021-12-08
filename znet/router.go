package znet

import "zinx/ziface"

// 先继承BaseRouter基类再重写具体的方法，这个有点像java中的抽象类
// 是因为有的Handle不希望有PreHandel和PostHandle这两个业务，所以router继承baserouter就可以不用实现全部方法，只需按需实现即可
type BaseRouter struct{}

func (br *BaseRouter) PreHandle(request ziface.IRequest) {}

func (br *BaseRouter) Handle(request ziface.IRequest) {}

func (br *BaseRouter) PostHandle(request ziface.IRequest) {}
