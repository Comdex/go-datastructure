package ds

type Queue interface {
	Clear()
	Length() int
	IsEmpty() bool
	Display()
	// 读取队首元素
	Peek() interface{}
	// 入队
	Offer(interface{}) error
	// 出队
	Poll() interface{}
}