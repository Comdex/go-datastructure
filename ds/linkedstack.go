package ds

import (
	"fmt"
)

type LinkedStack struct {
	top *Node
}
func NewLinkedStack() Stack {
	return &LinkedStack{nil}
}

func (this *LinkedStack) Clear() {
	this.top = nil
}

func (this *LinkedStack) Length() int {
	node := this.top
	length := 0
	for node != nil {
		node = node.GetNext()
		length++
	} 
	return length
}

func (this *LinkedStack) IsEmpty() bool {
	return this.top == nil
}

func (this *LinkedStack) Peek() interface{} {
	return this.top.GetData()
}

func (this *LinkedStack) Push(data interface{}) error {
	if this.top == nil {
		this.top = NewNode(data, nil)
	}else{
		node := NewNode(data,this.top)
		this.top = node
	}
	return nil
}

func (this *LinkedStack) Pop() interface{} {
	if this.top == nil {
		return nil
	}
	node := this.top
	this.top = this.top.GetNext()
	return node.GetData()
}

func (this *LinkedStack) Display() {
	node := this.top
	fmt.Println("LinkedStack Elem Display :")
	for node != nil {
		fmt.Printf("%v ",node.GetData())
		node = node.GetNext()
	}
	fmt.Println()
}
