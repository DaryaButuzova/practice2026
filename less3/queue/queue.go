package queue

type Queue[T any] struct {
	s         []T
	low, high int
	size      int
}

func NewQueue[T any](size int) *Queue[T] {
	return &Queue[T]{
		s:    make([]T, size),
		size: size,
		low:  -1,
		high: -1,
	}
}

func (q *Queue[T]) Push(v T) {

	if q.low == -1 {
		q.low = 0
		q.high = 0
		q.s[0] = v
		return
	}

	q.high = (q.high + 1) % q.size
	q.s[q.high] = v
}

// Pop - получение значения из очереди
func (q *Queue[T]) Pop() T {

	if q.low == -1 {
		panic("очередь пуста")
	}

	val := q.s[q.low]

	if q.low == q.high {
		q.low = -1
		q.high = -1
	} else {
		q.low = (q.low + 1) % q.size
	}

	return val
}
