package ds

type Stack interface {
	Clear()
	Length() int
	IsEmpty() bool
	Display()
	// 取栈顶值
	Peek() interface{}
	// 入栈
	Push(interface{}) error
	// 弹栈
	Pop() interface{}
}