package ds

import(
	"testing"
)

var queue Queue
func init() {
	queue = NewLinkedQueue()
	queue.Offer(56)
	queue.Offer(88)
	queue.Offer(99)
}

func Test_LinkedQueueOfferAndPeek(t *testing.T) {
	queue.Offer(299)
	queue.Display()
	
	if queue.Peek() != 56 {
		t.Error("LinkedQueue offer or peek error")
	}
}

func Test_LinkedQueuePollAndLength(t *testing.T) {
	if queue.Poll() != 56 {
		t.Error("LinkedQueue poll data error")
	}
	
	if queue.Peek() != 88 {
		t.Error("LinkedQueue poll error")
	}
	
	if queue.Length() != 3 {
		t.Error("LinkedQueue Length error")
	}
}