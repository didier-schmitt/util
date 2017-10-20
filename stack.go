package util

import "errors"

type stackItem struct {
	next *stackItem
	item *interface{}
}

// Stack is a LIFO (last in, first out) data collection.
// It can store any type of data.
type Stack struct {
	size int
	top  *stackItem
}

// Size returns the number of elements in the stack.
func (s *Stack) Size() int {
	return s.size
}

// Empty returns true if the stack is empty, false otherwise.
func (s *Stack) Empty() bool {
	return s.size == 0
}

// Push adds the given element on top of the stack.
func (s *Stack) Push(element interface{}) {
	var si = stackItem{
		next: s.top,
		item: &element,
	}
	s.top = &si
	s.size++
}

// Pop removes and returns the element on top of the stack.
func (s *Stack) Pop() (interface{}, error) {
	if s.size == 0 {
		return nil, errors.New("Cannot pop an empty stack")
	}
	var element = s.top.item
	s.top = s.top.next
	s.size--
	return *element, nil
}

// Peek returns the element on the top of the stack.
func (s *Stack) Peek() (interface{}, error) {
	if s.size == 0 {
		return nil, errors.New("Cannot peek an empty stack")
	}
	return *s.top.item, nil
}
