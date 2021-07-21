package main

import "fmt"

func main() {
	//变量定义
	var name int
	name = 2
	fmt.Printf("让我试试%d", name)
	var num = 26
	fmt.Printf("他是%T   %d\n", num, num)

	//多个变量同时定义
	var a, b, c int
	a = 1
	b = 2
	c = 3
	fmt.Println(a, b, c)
	var m, n int = 100, 200
	fmt.Println(m, n)
	//同时多种可用类型推断
	var n1, f1, s1 = 1, 2.34, "1asdasd"
	fmt.Println(n1, f1, s1)

}
