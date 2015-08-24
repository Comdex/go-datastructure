package ds

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


