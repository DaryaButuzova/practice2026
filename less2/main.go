package main

import (
	"fmt"

	"gitlab.com/DaryaButuzova/practice2026/less2/stack"
)

func main() {

	s := stack.NewStack(5)

	stack.Push(s, 10)
	stack.Push(s, 20)
	stack.Push(s, 30)

	fmt.Println("Верхний элемент:", stack.Peek(s)) // Должно быть 30

	fmt.Println("Извлекли:", stack.Pop(s)) // 30
	fmt.Println("Извлекли:", stack.Pop(s)) // 20

	fmt.Println("Верхний элемент после извлечения:", stack.Peek(s)) // 10
}
