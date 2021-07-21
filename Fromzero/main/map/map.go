package main

import "fmt"

func main() {
	var map1 map[int]string
	var map2 = make(map[int]string)
	var map3 = map[int]int{1: 1, 2: 2, 3: 3}
	fmt.Println(map1, map2, map3)
	fmt.Println(map1 == nil, map2 == nil, map3 == nil)
	map2[0] = "你好"
	map2[1] = "????"
	map2[2] = "hello"
	map2[3] = "gash"
	map2[4] = "english"
	fmt.Println(map1, map2, map3)

	m, n := map2[4]
	if n{
		fmt.Println(m)
	}
	map2[3]="lll"
	fmt.Println(map2)
	delete(map2,3)
	fmt.Println(map2)
}
