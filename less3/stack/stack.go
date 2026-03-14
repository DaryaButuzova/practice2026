package stack

type Stack[T any] struct {
	s    []T
	head int
}

func NewStack[T any](size int) *Stack[T] {
	return &Stack[T]{
		s:    make([]T, size),
		head: -1,
	}
}

// Push - добавление в стек
func (s *Stack[T]) Push(v T) {
	s.head++
	s.s[s.head] = v
}

// Pop - получение и удаление вершины
func (s *Stack[T]) Pop() T {
	val := s.s[s.head]
	s.head--
	return val
}

// Peek - просмотр вершины
func (s *Stack[T]) Peek() T {
	return s.s[s.head]
}
