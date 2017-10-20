package util

import "testing"

func Test_Queue(t *testing.T) {
	var q = Queue{}
	if q.Size() != 0 {
		t.Error("Size should return 0 with an empty queue")
	}
	if !q.Empty() {
		t.Error("Empty should return true with an empty queue")
	}

	var x = "test"
	q.Enqueue(x)
	if q.Size() != 1 {
		t.Error("Enqueue should increase the stack size")
	}

	var n = 5
	for i := 0; i < n-1; i++ {
		q.Enqueue(i)
	}
	if q.Size() != n {
		t.Errorf("Size is wrong: expected %d but got %d", n, q.Size())
	}
	if q.Empty() {
		t.Error("Empty should be false with a non-empty queue")
	}

	var y, err = q.Peek()
	if err != nil || y == nil {
		t.Fatal("Peek should not fail on a non-empty queue")
	}
	if x != y {
		t.Error("Peek should return the first enqueued element")
	}
	if q.Size() != n {
		t.Error("Peek should not alter the queue size")
	}

	y, err = q.Dequeue()
	n--
	if err != nil || y == nil {
		t.Fatal("Dequeue should not fail on a non-empty queue")
	}
	if x != y {
		t.Error("Dequeue should return the first queued element")
	}
	if q.Size() != n {
		t.Error("Dequeue should decrease the queue size")
	}
}
