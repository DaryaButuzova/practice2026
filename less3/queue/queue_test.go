package queue

import "testing"

func TestQueue(t *testing.T) {

	q := NewQueue[int](3)

	q.Push(1)
	q.Push(2)

	if q.Pop() != 1 {
		t.Errorf("ожидали 1")
	}

	if q.Pop() != 2 {
		t.Errorf("ожидали 2")
	}
}
