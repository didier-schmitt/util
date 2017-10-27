package util

import (
	"errors"
	"fmt"
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

func (l *LinkedList) indexAt(index int) (*linkedListItem, error) {
	if l.size == 0 || index < 0 || index >= l.size {
		return nil, errors.New("index out of bounds")
	}

	var element = l.head
	for i := 0; i < index; i++ {
		element = element.next
	}
	return element, nil
}

// ElementAt returns the element at the given index.
// Fails if the index is greater than the list size.
// First index starts at 0.
func (l *LinkedList) ElementAt(index int) (interface{}, error) {
	var item, err = l.indexAt(index)
	if err != nil {
		return nil, err
	}
	return *item.item, nil
}

// InsertAt adds the given element next to the given index.
// See also: ElementAt(index).
func (l *LinkedList) InsertAt(element interface{}, index int) error {
	if index == l.size {
		l.Append(element)
		return nil
	}
	var item, err = l.indexAt(index)
	if err != nil {
		return err
	}

	var newItem = linkedListItem{
		next: item,
		item: &element,
	}

	var prevItem = item.prev
	if prevItem != nil {
		prevItem.next = &newItem
		newItem.prev = prevItem
	}
	item.prev = &newItem

	return nil
}

// RemoveAt removes the element at the given index.
// See also: ElementAt(index)
func (l *LinkedList) RemoveAt(index int) error {
	var item, err = l.indexAt(index)
	if err != nil {
		return err
	}
	var prevItem = item.prev
	var nextItem = item.next
	if prevItem != nil {
		prevItem.next = item.next
	}
	if nextItem != nil {
		nextItem.prev = item.prev
	}
	return nil
}

func (l *LinkedList) String() string {
	var sb string
	var s string
	var item = l.head
	for i := 0; i < l.size; i++ {
		s = fmt.Sprint(*item.item)
		if i > 0 {
			sb += ", "
		}
		sb += s
		item = item.next
	}
	return fmt.Sprintf("[%v]", sb)
}
