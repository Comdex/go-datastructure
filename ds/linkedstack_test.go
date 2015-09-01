package ds

import(
	"testing"
)

var stack Stack

func init() {
	stack = NewLinkedStack()
	stack.Push(56)
	stack.Push(79)
	stack.Push(100)
}

func Test_LinkedStackPushAndPeek(t *testing.T) {
	stack.Push(99)
	stack.Display()
	
	if stack.Peek() != 99 {
		t.Error("LinkedStack push or peek error")
	}
}

func Test_LinkedStackPop(t *testing.T) {
	if stack.Pop() != 99 {
		t.Error("LinkedStack pop data error")
	}
	
	if stack.Peek() != 100 {
		t.Error("LinkedStack pop error")
	}
}

func Test_LinkedStackLength(t *testing.T) {
	if stack.Length() != 3 {
		t.Error("LinkedStack length error")
	}
}

//func Test_LinkedStackClearAndEmpty(t *testing.T) {
//	stack.Clear()
//	if stack.IsEmpty() == false {
//		t.Error("LinkedStack Clear or IsEmpty error")
//	}
//}