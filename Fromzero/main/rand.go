package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	num := 0
	for i := 1; i < 10; i++ {
		num = rand.Intn(64)
		fmt.Println(num)
	}
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < 10; i++ {
		num = rand.Intn(64)
		fmt.Println(num)
	}

}
