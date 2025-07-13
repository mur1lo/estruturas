package main

import (
	"fmt"

	"github.com/mur1lo/estruturas/stack"
)

func main() {
	stack := stack.NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println("Stack items:", stack.Items())
	item, ok := stack.Peek()
	fmt.Println("Top item:", item, "Exists:", ok)
	fmt.Println("Stack size:", stack.Size())

	item, ok = stack.Pop()
	fmt.Println("Pooped:", item, "Exists:", ok)
	fmt.Println("Stack after pop:", stack.Items())
}
