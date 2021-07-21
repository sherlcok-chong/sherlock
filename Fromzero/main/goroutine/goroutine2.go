package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go printfA(&wg)
	go printfNum(&wg)
	wg.Wait()
}

func printfNum(wg *sync.WaitGroup)  {
	for i:=1;i<6;i++ {
		time.Sleep(250*time.Millisecond)
		fmt.Printf("%d",i)
	}
	wg.Done()
}

func printfA(wg *sync.WaitGroup){
	for i:='a';i<='e';i++ {
		time.Sleep(400*time.Millisecond)
		fmt.Printf("%c",i)
	}
	wg.Done()
}
