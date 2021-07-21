package main

import (
	"fmt"
	"os"
)

func main() {
	s := make([]byte, 2048)
	n, _ := os.Stdin.Read(s)
	fmt.Println(s[:n-1])
	s1 := string(s[:n-1])
	fmt.Println(s1)
	for i, val := range s {

		s[i] = val % 10+110
	}
	fmt.Println(s[:n-1])
	s1 = string(s[:n-1])
	fmt.Println(s1)
}
