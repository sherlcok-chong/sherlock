package main

import (
	"fmt"
)

func main() {
	//var a chan int
	//fmt.Printf("%T,%v\n",a,a)

	a := make(chan int)
	//fmt.Printf("%T,%v", a, a)

	go func() {
		for i := 1; i < 10; i++ {
			fmt.Println(i)
		}
		a <- 10
	}()
	data := <-a
	fmt.Println("over", data)
}
