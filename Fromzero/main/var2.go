package main

import "fmt"

func main() {
	var num = 100
	fmt.Printf("%d  %p\n", num, &num)
	const (
		a int = 100
		b
		c
		//如若未重新定义，则变量的类型还有值和上一个相同
	)
	fmt.Println(a, b, c)
	const (
		m = iota
		n
		k = 132
		p = iota
	)
	fmt.Println(m, n, p)
	var A bool = true
	B := false
	fmt.Printf("%T %T %t %t", A, B, A, B)

}
