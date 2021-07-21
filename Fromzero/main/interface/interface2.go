package main

import "fmt"

func main() {
	var a1 A
	a1 = human{12}
	var a2 A = animal{"cat"}
	test2(a1)
	test2(a2)

	//利用空接口实现可以储存各种数据的map和slice
	map1 := make(map[interface{}]interface{})
	map1[123] = "123"
	map1["123"] = 123
	map1[human{12}] = "nih"

	slice1 :=make([]interface{},0,10)
	slice1=append(slice1,"12","cox",a1,a2)

	fmt.Println(map1,slice1)
	
}

type A interface {
}
type human struct {
	age int
}

type animal struct {
	name string
}

func test2(a A) {
	fmt.Println(a)
}
