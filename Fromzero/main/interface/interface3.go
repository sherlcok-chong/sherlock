package main

import "fmt"

func main() {
	t := S{"li"}
	t.test1()
	t.test2()
	t.test3()
	check2(t)
	//fmt.Sprintf()
}

type t1 interface {
	test1()
}
type t2 interface {
	test2()
}
type all interface {
	t1
	t2
	test3()
}
type S struct {
	name string
}

func (a S) test1(){
	fmt.Println("test1")
}
func (a S) test2() {
	fmt.Println("test2")
}
func (a S) test3()  {
	fmt.Println("test3")
}

func check1(a t1)  {
	a.test1()
}
func check2(a t2){
	a.test2()
}
func check3(a all)  {
	a.test1()
	a.test2()
	a.test3()
}
