package ds

import (
	"fmt"
	"errors"
)

// 线性表顺序表实现
type SqList struct {
	data []interface{}
}

func NewSqList() List {
	var list List = new(SqList)
	return list
}

func (this *SqList) Length() int {
	return len(this.data)
}

func (this *SqList) Clear() {
	this.data = nil
}

func (this *SqList) IsEmpty() bool {
	return len(this.data) == 0
}

func (this *SqList) Get(index int) (data interface{}, err error) {
	if index < 0 || index >= this.Length() {
		err = errors.New("The index is illegal!")
		return
	}
	
	data = this.data[index]
	return
}

func (this *SqList) Insert(index int, data interface{}) error {
	if index < 0 || index > this.Length() {
		return errors.New("The index is illegal!")
	}
	var temp []interface{} = make([]interface{}, len(this.data[index:]))
	copy(temp,this.data[index:])
	current := append(this.data[:index],data)
	this.data = append(current,temp...)
	return nil
}

func (this *SqList) Remove(index int) error {
	if index < 0 || index > this.Length() {
		return errors.New("The index is illegal!")
	}
	this.data = append(this.data[:index],this.data[index+1:]...)
	return nil
}

func (this *SqList) IndexOf(data interface{}) int {
	for i, val := range this.data {
		if val == data {
			return i
		}
	}
	return -1
}

func (this *SqList) Display() {
	for _, val := range this.data {
		fmt.Printf("%v ",val)
	}
	fmt.Println()
}