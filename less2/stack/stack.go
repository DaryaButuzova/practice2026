package stack

type stack struct {
	s    []any // слайс в котором хранятся значения в стеке
	head int   // индекс головы стека
}

func NewStack(size int) *stack {
	return &stack{
		s:    make([]any, size),
		head: -1,
	}
}

// push - добавление в стек значения
func Push(s *stack, v any) {
	s.head++
	s.s[s.head] = v
}

// pop - получения значения из стека и его удаление из вершины
func Pop(s *stack) any {
	val := s.s[s.head]
	s.head--
	return val
}

// peek - просмотр значения на вершине стека
func Peek(s *stack) any {
	return s.s[s.head]
}
