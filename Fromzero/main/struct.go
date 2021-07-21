package main

import "fmt"

type student struct {
	name string
	id string
	class string
}
func main() {
	var a student
	a.id="04201204"
	a.name="sherlock"
	a.class="2007"
	fmt.Println(a)

	b := student{}
	b.id="170106"
	b.name="wangling"
	b.class="three"
	fmt.Println(b)

	c:=student{
		name: "change",
		id: "1122",
		class: "six",
	}
	fmt.Println(c)

	d:=student{"123","123","345"}
	fmt.Println(d)

	e:= struct {
		name string
		age int
	}{
		name: "种",
		age: 18,
	}
	fmt.Println(e)

}

type ni struct {
	string
	int
	//匿名字段缺陷很大
	//把类型名字作为变量名
}
