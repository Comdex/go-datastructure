package ds

type Node struct {
	data interface{} //结点值
	next *Node //后继结点
}

func NewNode(data interface{}, next *Node) *Node {
	return &Node{data:data, next:next}
}

func (this *Node) SetData(data interface{}) {
	this.data = data
}

func (this *Node) GetData() interface{} {
	return this.data
}

func (this *Node) SetNext(next *Node) {
	this.next = next
}

func (this *Node) GetNext() *Node {
	return this.next
}