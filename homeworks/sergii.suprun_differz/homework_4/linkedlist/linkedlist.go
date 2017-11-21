package linkedlist

import "errors"
import "fmt"

// LinkedList data structure
type LinkedList struct {
	head *node
	last *node
	size int
}

// New creates new list
func New() *LinkedList {
	return &LinkedList{}
}

// Add element to list
func (t *LinkedList) Add(e interface{}) {
	newNode := newNode(e, nil)

	if t.last == nil {
		t.head = newNode
	} else {
		t.last.setNext(newNode)
	}
	t.last = newNode
	t.size++
}

// Get element by index
func (t *LinkedList) Get(index int) (interface{}, error) {
	error := t.checkIndex(index)
	if error != nil {
		return nil, error
	}
	node := t.head
	for i := 0; i < index; i++ {
		node = node.getNext()
	}
	return node.getValue(), nil
}

// Remove implements method overloading (by value or by index)
func (t *LinkedList) Remove(args interface{}) bool {
	switch v := args.(type) {
	case int:
		return t.removeIndex(v)
	default:
		return t.remove(v)
	}
}

func (t *LinkedList) remove(e interface{}) bool {
	prev := t.head
	for node := t.head; node != nil; node = node.getNext() {
		value := node.getValue()
		if (e == nil) && (value == nil) || (e != nil) && (e == value) {
			t.unlink(prev, node)
			return true
		}
		prev = node
	}
	return false
}

func (t *LinkedList) removeIndex(index int) bool {
	error := t.checkIndex(index)
	if error != nil {
		return false
	}
	var prev *node
	node := t.head
	for i := 0; i < index; i++ {
		prev = node
		node = node.getNext()
	}
	t.unlink(prev, node)
	return true
}

func (t *LinkedList) unlink(prev *node, node *node) {
	next := node.getNext()
	if prev == nil {
		t.head = next
	} else {
		prev.setNext(next)
	}
	if next == nil {
		t.last = prev
	}
	t.size--
}

func (t *LinkedList) checkIndex(index int) error {
	if (index >= t.size) || (index < 0) {
		return errors.New("index is out of bounds")
	}
	return nil
}

// IsEmpty returns true if list is empty
func (t *LinkedList) IsEmpty() bool {
	return t.size == 0
}

// Size of list
func (t *LinkedList) Size() int {
	return t.size
}

func (t *LinkedList) String() string {
	s := "["
	for node := t.head; node != nil; node = node.getNext() {
		if node != t.head {
			s += ", "
		}
		s += fmt.Sprint(node.getValue())
	}
	return s + "]"
}
