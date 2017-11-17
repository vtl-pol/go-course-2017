package list

// List is the common interface implemented by LinkedList
type List interface {
	Add(e interface{})
	Remove(e interface{}) bool
	Set(index uint64, e interface{})
	Get(index uint64) interface{}
	Size() uint64
	IsEmpty() uint64
}
