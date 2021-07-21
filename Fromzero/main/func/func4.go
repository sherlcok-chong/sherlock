package main

import "fmt"

func main() {
	f:=into
	fmt.Println(f)
	f2:=into()
	fmt.Println(f2)
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f2())

	f3:=into()
	fmt.Println(f3())
	fmt.Println(f3())

}

func into() func()int {
	i:=0
	fun:=func ()int{
		i++
		return i
	}
	return fun
}