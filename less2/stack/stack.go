package stack

type Stack struct {
	s    []any // слайс в котором хранятся значения в стеке
	head int   // индекс головы стека
}

func NewStack(size int) *Stack {
	return &Stack{
		s:    make([]any, size),
		head: -1,
	}
}

// push - добавление в стек значения
func Push(s *Stack, v any) {
	s.head++
	s.s[s.head] = v
}

// pop - получения значения из стека и его удаление из вершины
func Pop(s *Stack) any {
	val := s.s[s.head]
	s.head--
	return val
}

// peek - просмотр значения на вершине стека
func Peek(s *Stack) any {
	return s.s[s.head]
}
