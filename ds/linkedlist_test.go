package ds

import (
	"testing"
)

func Test_LinkedListInsert(t *testing.T) {
	list := NewLinkedList2()
	list.Insert(0, 55)
	list.Insert(1, 66)
	list.Insert(2, 77)

	if list.IndexOf(55) != 0 || list.IndexOf(66) != 1 || list.IndexOf(77) != 2 {
		t.Error("LinkedList Insert error")
	}
}

func Test_LinkedListReverse(t *testing.T) {
	list := NewLinkedList2()
	list.Insert(0, 55)
	list.Insert(1, 66)
	list.Insert(2, 77)
	list.Reverse()

	if list.IndexOf(55) != 2 || list.IndexOf(66) != 1 || list.IndexOf(77) != 0 {
		t.Error("LinkedList Reverse error")
	}
}
