package ds

import (
	"fmt"
)

type LinkedQueue struct {
	head *Node
}
func NewLinkedQueue() Queue {
	return &LinkedQueue{nil}
}

func (this *LinkedQueue) Clear() {
	this.head = nil
}

func (this *LinkedQueue) Length() int {
	node := this.head
	length := 0 
	for node != nil {
		node = node.GetNext()
		length++
	}
	return length
}

func (this *LinkedQueue) Display() {
	node := this.head
	fmt.Println("LinkedQueue Elem Display :")
	for node != nil {
		fmt.Printf("%v ", node.GetData())
		node = node.GetNext()
	}
	fmt.Println()
}

func (this *LinkedQueue) IsEmpty() bool {
	return this.head == nil
}

func (this *LinkedQueue) Offer(data interface{}) error {
	if this.head == nil {
		this.head = NewNode(data, nil)
		return nil
	}
	
	node := this.head
	for node.GetNext() != nil {
		node = node.GetNext()
	}
	node.SetNext(NewNode(data, nil))
	return nil
}

func (this *LinkedQueue) Peek() interface{} {
	return this.head.GetData()
}

func (this *LinkedQueue) Poll() interface{} {
	node := this.head
	this.head = this.head.GetNext()
	return node.GetData()
}