package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	go func() {
		time.Sleep(1000*time.Millisecond)
		ch1 <- 1
		fmt.Println("执行成功")
	}()
	select {
	case <-ch1:
		fmt.Println("执行了ch1")
	case <-ch2:
		fmt.Println("执行了ch2")
	case <-ch3:
		fmt.Println("执行了ch3")
	case <-time.After(5*time.Millisecond):
		fmt.Println("123")
	//default:
	//	fmt.Println("执行default")
	}

}
