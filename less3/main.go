package main

import (
	"fmt"

	"github.com/DaryaButuzova/practice2026/less3/queue"
	"github.com/DaryaButuzova/practice2026/less3/stack"
	"github.com/DaryaButuzova/practice2026/less3/tree"
)

func main() {

	// стек
	st := stack.NewStack[int](5)
	st.Push(10)
	st.Push(20)
	fmt.Println(st.Pop())

	// очередь
	q := queue.NewQueue[int](5)
	q.Push(1)
	q.Push(2)
	fmt.Println(q.Pop())

	// дерево
	t := tree.NewTree[int]()
	t.Add(5)
	t.Add(2)
	t.Add(8)

	fmt.Println(t.Values())
}
