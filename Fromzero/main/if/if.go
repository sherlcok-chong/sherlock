package main

import "fmt"

func main() {
	var num int
	if fmt.Scanln(&num);num > 33 {
		fmt.Printf("%d\n", num)
	} else if num > 31 {
		fmt.Printf("OK\n")
	} else {
		fmt.Printf("12\n")
	}
	fmt.Println(num)
}
