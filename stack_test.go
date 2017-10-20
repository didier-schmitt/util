package util

import "testing"

func Test_Stack(t *testing.T) {
	var s = Stack{}
	if s.Size() != 0 {
		t.Error("A new Stack should have a size of zero")
	}
	if !s.Empty() {
		t.Error("Empty should be true with an empty stack")
	}

	var n = 5
	for i := 0; i < n; i++ {
		s.Push(i)
	}
	if s.Size() != n {
		t.Errorf("Size is wrong: expected %d but got %d", s.Size(), n)
	}
	if s.Empty() {
		t.Error("Empty should be false with a non-empty stack")
	}

	var x = "test"
	s.Push(x)
	n++
	if s.Size() != n {
		t.Error("Push should increase the stack size")
	}

	var y, err = s.Peek()
	if err != nil || y == nil {
		t.Fatal("Peek should not fail on a non-empty stack")
	}
	if x != y {
		t.Error("Peek should return the last pushed element")
	}
	if s.Size() != n {
		t.Error("Peek should not alter the stack size")
	}

	y, err = s.Pop()
	n--
	if err != nil || y == nil {
		t.Fatal("Pop should not fail on a non-empty stack")
	}
	if x != y {
		t.Error("Pop should return the last pushed element")
	}
	if s.Size() != n {
		t.Error("Pop should decrease the stack size")
	}
}
