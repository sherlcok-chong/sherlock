package main

import (
	"fmt"
	"time"
)

func main() {
	go printNum()
	for i :=1;i<100;i++{
		fmt.Println("A")
	}
	time.Sleep(1*time.Second)

}

func printNum(){
	for i :=1;i<100;i++{
		fmt.Println(i)
	}
	return
}
func testes(){
	go printNum()
}