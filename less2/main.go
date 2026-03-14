package main

import (
	"fmt"

	"github.com/DaryaButuzova/practice2026/less2/queue"
	"github.com/DaryaButuzova/practice2026/less2/stack"
)

func main() {

	// стек
	st := stack.NewStack(5)

	stack.Push(st, 10)
	stack.Push(st, 20)

	fmt.Println("stack pop:", stack.Pop(st))
	fmt.Println("stack peek:", stack.Peek(st))

	// очередь
	q := queue.NewQueue(5)

	queue.Push(q, 1)
	queue.Push(q, 2)

	fmt.Println("queue pop:", queue.Pop(q))
}
