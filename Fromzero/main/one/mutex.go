package main

import (
	"fmt"
"math/rand"
"sync"
"time"
)

var ticket2 int = 10
var mutex2 sync.Mutex
func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	go saleTicket2("1号",&wg)
	go saleTicket2("2号",&wg)
	go saleTicket2("3号",&wg)
	go saleTicket2("4号",&wg)
	wg.Wait()
	//time.Sleep(1*time.Second)
}

func saleTicket2(name string,wg *sync.WaitGroup) {
	rand.Seed(time.Now().UnixNano())
	for {
		mutex2.Lock()
		if ticket2 > 0 {
			time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)
			fmt.Println(name, "售出", ticket2)
			ticket2--
		} else {
			mutex2.Unlock()
			fmt.Println(name,"无了")
			break
		}
		mutex2.Unlock()
	}
	wg.Done()
}
