package main

func main() {
	st := NewStack(5)
	Push(st, 10)
	fmt.println(Pop(st))
}
