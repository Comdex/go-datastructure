package ds

import (
	"strconv"
	"os"
	"bufio"
	"errors"
	"fmt"
)

//链表结点
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

// 线性表链表实现
type LinkedList struct {
	head *Node
}

func NewLinkedList(n int, order bool) (list List) {
	scanner := bufio.NewScanner(os.Stdin)
	node := NewNode(nil,nil)
	list  = &LinkedList{node}
	if order {
		fmt.Println("please input:")
		for i:=0; i<n; i++ {
			if scanner.Scan() {
				j , _ := strconv.Atoi(scanner.Text())
				list.Insert(0,j)
			}
		}
		return list
	}else {
		fmt.Println("please input:")
		for i:=0; i<n; i++ {
			if scanner.Scan() {
				j, _ := strconv.Atoi(scanner.Text())
				list.Insert(list.Length(),j)
			}
		}
		return list
	}
}

func NewLinkedList2() *LinkedList {
	node := NewNode(nil,nil)
	return &LinkedList{node}
}

func (this *LinkedList) Length() int {
	p := this.head.GetNext()
	length := 0
	for p != nil {
		p = p.GetNext()
		length++
	}
	return length
}

func (this *LinkedList) Clear() {
	this.head.SetData(nil)
	this.head.SetNext(nil)
}

func (this *LinkedList) IsEmpty() bool {
	return this.head.GetNext() == nil
}

func (this *LinkedList) Get(index int) (interface{}, error) {
	if index < 0 || index >= this.Length() {
		return nil, errors.New("The index is illegal!")
	}
	p := this.head.GetNext()
	location := 0
	for p != nil && location < index {
		p = p.GetNext()
		location ++
	}
	return p.GetData() , nil
}

func (this *LinkedList) Insert(index int, data interface{}) error {
	if index < 0 || index > this.Length() {
		return errors.New("The index is illegal!")
	}
	p := this.head
	location := -1
	for p != nil && location < (index - 1) {
		p = p.GetNext()
		location++
	}
	if p == nil {
		return errors.New("The index is illegal!")
	}
	node := NewNode(data, nil)
	node.SetNext(p.GetNext())
	p.SetNext(node)
	return nil
}

func (this *LinkedList) Remove(index int) error {
	if index < 0 || index >= this.Length() {
		return errors.New("The index is illegal!")
	}
	p := this.head
	location := -1
	for p != nil && location < index-1 {
		p = p.GetNext()
		location ++
	}
	
	if p == nil  {
		return errors.New("The index is illegal!")
	}
	p.SetNext(p.GetNext().GetNext())
	return nil
}

func (this *LinkedList) IndexOf(data interface{}) int {
	p := this.head.GetNext()
	location := 0
	for p != nil && p.GetData() != data {
		p = p.GetNext()
		location++
	}
	if p != nil {
		return location
	}else{
		return -1
	}
}

func (this *LinkedList) Display() {
	p := this.head
	for p.GetNext() != nil {
		p = p.GetNext()
		fmt.Printf("%v ",p.GetData())
	}
	fmt.Println()
}

// 反转链表
func (this *LinkedList) Reverse() {
	// 第一个结点
	first := this.head.GetNext()
	var prev *Node
	// 反转完成后第一个结点的下一个结点就是nil
	for first != nil {
		next := first.GetNext()
		first.SetNext(prev)
		prev = first
		first = next
	}
	this.head.SetNext(prev)
}