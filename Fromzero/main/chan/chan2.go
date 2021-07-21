package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg = sync.WaitGroup{}
	wg.Add(2)
	a := make(chan int)
	go APrint(a, &wg)
	go numPrint(a, &wg)
	wg.Wait()
}

func APrint(a chan int, wg *sync.WaitGroup) {
	for i := 'a'; i < 'j'; i++ {
		<-a
		fmt.Println(string(i))
		a <- int(i)
	}
	wg.Done()
}

func numPrint(a chan int, wg *sync.WaitGroup) {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
		a <- i
		<-a
	}
	wg.Done()
}
