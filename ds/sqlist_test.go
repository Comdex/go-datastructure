package ds

import (
	"testing"
)

var list List

func init() {
	list = NewSqList()
	list.Insert(0, 35)
	list.Insert(1, 89)
	list.Insert(2, 25)
	list.Insert(3, "hi")
}

func Test_SqListInsertAndIndex(t *testing.T) {

	if list.IndexOf(35) != 0 || list.IndexOf(89) != 1 || list.IndexOf(25) != 2 || list.IndexOf("hi") != 3 {
		t.Error("SqList Insert error")
	}
}

func Test_SqListGet(t *testing.T) {
	data, _ := list.Get(3)
	if data != "hi" {
		t.Error("SqList Get error")
	}
}

func Test_SqListRemove(t *testing.T) {
	list.Remove(1)
	if data, _ := list.Get(1); data == 89 {
		t.Error("SqList Remove error")
	}
}

func Test_SqListLength(t *testing.T) {
	if list.Length() != 3 {
		t.Error("SqList Length error")
	}
}

func Test_SqListClear(t *testing.T) {
	list.Clear()

	if list.Length() != 0 {
		t.Error("SqList Clear error")
	}
}
