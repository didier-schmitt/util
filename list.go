package util

import (
	"errors"
)

type linkedListItem struct {
	prev *linkedListItem
	next *linkedListItem
	item *interface{}
}

// LinkedList is double-linked list.
type LinkedList struct {
	size int
	head *linkedListItem
	tail *linkedListItem
}

// Size returns the number of elements in the list.
func (l *LinkedList) Size() int {
	return l.size
}

// Empty returns true if the list has no element, false otherwise.
func (l *LinkedList) Empty() bool {
	return l.size == 0
}

// Append adds the given element at the end of the list.
func (l *LinkedList) Append(element interface{}) {
	var li = linkedListItem{
		item: &element,
	}
	if l.tail != nil {
		l.tail.next = &li
		li.prev = l.tail
	}
	l.tail = &li
	if l.head == nil {
		l.head = &li
	}
	l.size++
}

// Tail returns the element at the end of the list.
func (l *LinkedList) Tail() (interface{}, error) {
	if l.size == 0 {
		return nil, errors.New("cannot tail an empty list")
	}
	return *l.tail.item, nil
}

// Head returns the element at the beginning of the list.
func (l *LinkedList) Head() (interface{}, error) {
	if l.size == 0 {
		return nil, errors.New("cannot head on an empty list")
	}
	return *l.head.item, nil
}

// Prepend adds the given element at the beginning of the list.
func (l *LinkedList) Prepend(element interface{}) {
	var li = linkedListItem{
		item: &element,
	}
	if l.head != nil {
		l.head.prev = &li
		li.next = l.head
	}
	l.head = &li
	if l.tail == nil {
		l.tail = &li
	}
	l.size++
}
