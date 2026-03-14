package queue

type Queue struct {
	s         []any
	low, high int
	size      int
}

func NewQueue(size int) *Queue {
	return &Queue{
		s:    make([]any, size),
		size: size,
		low:  -1,
		high: -1,
	}
}

func Push(q *Queue, v any) {
	if q.low == -1 {
		q.low = 0
		q.high = 0
		q.s[0] = v
	}
	q.high = (q.high + 1) % q.size
	q.s[q.high] = v

}

// pop - получения значения из очереди и его удаление
func Pop(q *Queue, v any) {
	if q.low == -1 {
		panic("очередь пуста")

	}
	if q.low == q.high {
		q.low = -1
		q.high = -1
	} else {
		q.low = (q.low + 1) % q.size
	}

}
