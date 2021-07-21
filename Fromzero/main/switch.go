package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)
	switch num {
	case 1:
		fmt.Println("1")
		num = 2
	case 2:
		fmt.Println("3")
	default:
		fmt.Println("4")
	}
	fmt.Println(num)
}
