package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ticket int = 10
var mutex sync.Mutex
func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	go saleTicket("1号",&wg)
	go saleTicket("2号",&wg)
	go saleTicket("3号",&wg)
	go saleTicket("4号",&wg)
	wg.Wait()
	//time.Sleep(1*time.Second)
}

func saleTicket(name string,wg *sync.WaitGroup) {
	rand.Seed(time.Now().UnixNano())
	for {
		mutex.Lock()
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)
			fmt.Println(name, "售出", ticket)
			ticket--
		} else {
			mutex.Unlock()
			fmt.Println(name,"无了")
			break
		}
		mutex.Unlock()
	}
	wg.Done()
}
