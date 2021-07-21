package main

import "fmt"

func main() {
	f1 := mouse{"阿米洛"}
	//fmt.Println(f1.name)
	test(f1)
	
}

type USB interface {
	start()
	end()
}

type mouse struct {
	name string
}

func (m mouse) start(){
	fmt.Println("linkStart")
}

func (m mouse) end()  {
	fmt.Println("linkOver")
}

func test(p USB)  {
	p.start()
	p.end()
}