package main

import "fmt"

func main() {
	num := 0
	sum:= 1
	fmt.Scanln(&num)
	for i := 1; i <= num; i++ {
		sum*=i
	}
	fmt.Println(sum)
}
