package main

import (
	"fmt"
)

func main() {
	q := Queue{1, 2, 3, 4, 5}
	q, _ = q.Pop()
	fmt.Println(q)
	q = q.Push(0)
	fmt.Println(q)
	fmt.Println(q.Empty())
	fmt.Println(q.Head())
	fmt.Println(q.Tail())
}

type Queue []interface{}

func (q Queue) Pop() ([]interface{}, interface{}) {
	val := q[0]
	return q[1:], val
}
func (q Queue) Push(val interface{}) []interface{} {
	return append(q, val)
}
func (q Queue) Empty() bool {
	return !(len(q) > 0)
}

func (q Queue) Head() interface{} {
	return q[0]
}

func (q Queue) Tail() interface{} {
	return q[len(q)-1]
}
