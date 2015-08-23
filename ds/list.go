package ds

import (
	"strconv"
	"os"
	"bufio"
	"errors"
	"fmt"
)

type List interface {
	Length() int
	Clear()
	IsEmpty() bool
	Get(int) (interface{}, error)
	Insert(int , interface{}) error
	Remove(int) error
	IndexOf(interface{}) int
	Display()
}


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