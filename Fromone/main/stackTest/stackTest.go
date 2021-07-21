package main

import (
	"Fromone/stackAndqune/stack"
	"fmt"
)

func main() {
	s := stack.Stack{1,2,3,4}

	fmt.Println(s)
	s, _ = s.Pop()
	fmt.Println(s)
	s = s.Push(12)
	fmt.Println(s)
	fmt.Println(s.Empty())
	fmt.Println(s.Top())
}


