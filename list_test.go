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
}
