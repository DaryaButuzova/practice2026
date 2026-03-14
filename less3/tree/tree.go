package tree

type Tree[T int] struct {
	head *node[T]
}

type node[T int] struct {
	left, right *node[T]
	v           T
}

func NewTree[T int]() *Tree[T] {
	return &Tree[T]{}
}

// Add - добавление значения
func (t *Tree[T]) Add(v T) {

	if t.head == nil {
		t.head = &node[T]{v: v}
		return
	}

	current := t.head

	for {
		if v < current.v {
			if current.left == nil {
				current.left = &node[T]{v: v}
				return
			}
			current = current.left
		} else {
			if current.right == nil {
				current.right = &node[T]{v: v}
				return
			}
			current = current.right
		}
	}
}

// Values - получение отсортированных значений
func (t *Tree[T]) Values() []T {

	var result []T

	var walk func(n *node[T])

	walk = func(n *node[T]) {
		if n == nil {
			return
		}

		walk(n.left)
		result = append(result, n.v)
		walk(n.right)
	}

	walk(t.head)

	return result
}
