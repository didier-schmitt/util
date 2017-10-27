package util

import (
	"testing"
)

func Test_LinkedList(t *testing.T) {
	var l = LinkedList{}
	if l.Size() != 0 {
		t.Error("Size should return 0 with an empty list")
	}
	if !l.Empty() {
		t.Error("Empty should return true with an empty list")
	}

	var n = 5
	for i := 0; i < n; i++ {
		l.Append(i)
	}
	if l.Size() != n {
		t.Error("Append should increase the list size")
	}

	var x, err = l.Tail()
	if err != nil || x == nil {
		t.Error("Tail should not fail on a non-empty list")
	}
	if l.Size() != n {
		t.Error("Tail should not alter the list size")
	}
	if x != n-1 {
		t.Errorf("Wrong tail value: expected %d but got %d", n-1, x)
	}

	x, err = l.Head()
	if err != nil || x == nil {
		t.Error("Head should not fail on a non-empty list")
	}
	if l.Size() != n {
		t.Error("Head should not alter the list size")
	}
	if x != 0 {
		t.Errorf("Wrong head value: expected %d but got %d", 0, x)
	}

	var m = 10
	for i := n; i < m; i++ {
		l.Prepend(i)
	}
	if l.Size() != m {
		t.Error("Prepend should increase the list size")
	}
	x, err = l.Head()
	if x != m-1 {
		t.Error("Prepend should add element at beginning")
	}

	x, err = l.ElementAt(99)
	if err == nil || x != nil {
		t.Error("ElementAt should fail with an index greater than the list size")
	}
	x, err = l.ElementAt(-1)
	if err == nil || x != nil {
		t.Error("ElementAt should fail with a negative index")
	}
	var el = LinkedList{}
	x, err = el.ElementAt(0)
	if err == nil || x != nil {
		t.Error("ElementAt should fail with an empty list")
	}
	x, err = l.ElementAt(3)
	if x != 6 {
		t.Errorf("Wrong element at index 3 : expected 6 but got %d", x)
	}

	err = el.InsertAt(10, 1)
	if err == nil {
		t.Error("InsertAt should fail with an empty list and a non zero index")
	}
	err = l.InsertAt(10, l.Size())
	if err != nil {
		t.Error("InsertAt should not fail with an index equal to the list size")
	}
	x, _ = l.Tail()
	if x != 10 {
		t.Error("InsertAt should append the value to the end of the list")
	}
	err = l.InsertAt(11, 4)
	if err != nil {
		t.Error("InsertAt should not fail with an index within the list index bounds")
	}
	x, _ = l.ElementAt(4)
	if x != 11 {
		t.Error("InsertAt should add an element at the given index")
	}

	err = el.RemoveAt(0)
	if err == nil {
		t.Error("RemoveAt should fail with an empty list")
	}
	err = l.RemoveAt(4)
	if err != nil {
		t.Error("RemoveAt should not fail with an index in the list index bounds")
	}
	x, _ = l.ElementAt(4)
	if x == 11 {
		t.Error("RemoveAt should remove an element at the given index")
	}
}
