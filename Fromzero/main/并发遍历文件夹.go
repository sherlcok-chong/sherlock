package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var cnt int

var mutex8 sync.Mutex
func main() {
	dirname := "D:\\"
	var wg sync.WaitGroup
	start := time.Now()
	ch1 := make(chan int)
	wg.Add(1)
	go listFiles2(dirname, ch1, &wg)
	wg.Wait()
	ended := time.Now().Sub(start)
	fmt.Println(ended, cnt)
}

func listFiles2(dirname string, ch1 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fileInfos, err := os.ReadDir(dirname)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, i := range fileInfos {
		filename := dirname + "\\" + i.Name()
		//fmt.Println(filename)
		mutex8.Lock()
		cnt++
		mutex8.Unlock()
		if i.IsDir() {
			wg.Add(1)
			go listFiles2(filename, ch1, wg)
		}
	}

	return
}
