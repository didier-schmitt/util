package util

import (
	"errors"
)

type queueItem struct {
	next *queueItem
	item *interface{}
}

// Queue is a FIFO (first in, first out) data collection.
// it can store any type of data.
type Queue struct {
	size int
	head *queueItem
	tail *queueItem
}

// Size returns the number of elements in the queue.
func (q *Queue) Size() int {
	return q.size
}

// Empty returns true if the queue has no element, false otherwise.
func (q *Queue) Empty() bool {
	return q.size == 0
}

// Enqueue adds the given element at the end of the queue.
func (q *Queue) Enqueue(element interface{}) {
	var qi = queueItem{
		item: &element,
	}
	if q.tail != nil {
		q.tail.next = &qi
	}
	q.tail = &qi
	if q.head == nil {
		q.head = &qi
	}
	q.size++
}

// Peek returns the element at the beginning of the queue.
func (q *Queue) Peek() (interface{}, error) {
	if q.size == 0 {
		return nil, errors.New("cannot peek an empty queue")
	}
	return *q.head.item, nil
}

// Dequeue remove and returns the element at the beginning of the queue.
func (q *Queue) Dequeue() (interface{}, error) {
	if q.size == 0 {
		return nil, errors.New("cannot dequeue an empty queue")
	}
	var element = q.head.item
	q.head = q.head.next
	q.size--
	if q.size == 0 {
		q.tail = nil
	}
	return *element, nil
}
