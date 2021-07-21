package main

import "fmt"

func main() {
	var p1 person
	p1.name = "ChangZhao"
	p1.age = 18
	p1.printInform()
	p2 := new(Wei)
	p2.age=19
	p2.name="Wing"
	p2.printInform()
	fmt.Println(p2.name)
}

type person struct {
	name string
	age  int
}
type Wei struct {
	name string
	age  int
}

func (p person) printInform() {
	fmt.Println(p.name, p.age)
}
func (p *Wei) printInform() {
	fmt.Println(p.age, p.name)
	p.name="chowing"
}
