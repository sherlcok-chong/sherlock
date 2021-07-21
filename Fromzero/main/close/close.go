package main

import "fmt"

func main() {
	a := make(chan int)
	go sendData(a)
	for value := range a{
		fmt.Println(value)
	}
	fmt.Println("over")
}

func sendData(a chan int)  {
	for i:=1;i<10;i++ {
		a<-i
	}
	close(a)
}