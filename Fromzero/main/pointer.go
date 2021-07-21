package main

import "fmt"

func main() {
	arr := [4]int{1,2,3,4}
	p1:= &arr
	var p2 *[4]int
	p2 = &arr
	fmt.Println(p1,*p1,p2)
	a:=1
	b:=2
	c:=3
	arr2 := [3]int{a,b,c}
	arr3 := [3]*int{&a,&b,&c}
	fmt.Println(arr2,arr3)
	var p int
	p=999993147483647
	p+=10
	fmt.Println(p)
}
