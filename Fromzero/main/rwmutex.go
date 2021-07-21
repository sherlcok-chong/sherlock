package main

import (
	"fmt"
	"sync"
	"time"
)

//写锁定情况下，对读写锁进行读锁定或者写锁定，都将阻塞；而且读锁与写锁之间是互斥的；
//读锁定情况下，对读写锁进行写锁定，将阻塞；加读锁时不会阻塞；
//对未被写锁定的读写锁进行写解锁，会引发Panic；
//对未被读锁定的读写锁进行读解锁的时候也会引发Panic；
//写解锁在进行的同时会试图唤醒所有因进行读锁定而被阻塞的goroutine；
//读解锁在进行的时候则会试图唤醒一个因进行写锁定而被阻塞的goroutine。
func main() {
	var wg sync.WaitGroup
	var rwMutex sync.RWMutex
	wg.Add(4)
	go writeData(1,&wg,&rwMutex)
	go writeData(2,&wg,&rwMutex)
	//go writeData(2,&wg,&rwMutex)
	go readData(1,&wg,&rwMutex)
	go readData(2,&wg,&rwMutex)
	wg.Wait()
}

func writeData(i int, wg *sync.WaitGroup, rwMutex *sync.RWMutex) {
	defer wg.Done()
	fmt.Println(i,"开始写")
	rwMutex.Lock()//写上锁
	//time.Sleep(1*time.Second)
	fmt.Println(i,"正在写")
	time.Sleep(1*time.Second)
	rwMutex.Unlock()//写解锁
	fmt.Println(i,"写完了")
}

func readData(i int, wg *sync.WaitGroup, rwMutex *sync.RWMutex) {
	defer wg.Done()
	time.Sleep(2*time.Millisecond)
	fmt.Println(i,"开始读")
	rwMutex.RLock()//读上锁
	//time.Sleep(1*time.Second)
	fmt.Println(i,"正在读")
	time.Sleep(1*time.Second)
	rwMutex.RUnlock()//读解锁
	fmt.Println(i,"读完了")
}
