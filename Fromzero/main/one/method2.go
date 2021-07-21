package main

import "fmt"

func main() {
	s1 := Father{"hong",18}
	s2 := Son{Father{"hui",12},"XUPT"}
	fmt.Println(s1.age)
	s1.where()
	fmt.Println(s2.name)
	s2.where()

	s2.here()
}

type Father struct {
	name string
	age int
}
type Son struct {
	Father
	School string
}

func (p Father) where()  {
	fmt.Println("这是他爸")
}
func (p Son) where()  {
	fmt.Println("这是他儿")
}
func (p Son) here()  {
	fmt.Println("他儿在这")
}