package main

import (
	"fmt"
	. "package/Link"
)

type person struct {
	name string
	age int
}
func main() {
	var Head *Linklist = nil
	fmt.Println(Head)
	Head = NewEmpty(2, Head)
	PrintAllList(Head)
	AssignmentEmpty(Head, person{name: "asd",age: 123},person{"qwe",456})
	PrintAllList(Head)
	DeleteData(AddressFind("asd", Head), Head)
	PrintAllList(Head)
	Change(AddressFind(2, Head), 3)
	PrintAllList(Head)
	Add(AddressFind(1, Head))
	Add(AddressFind(3, Head))
	AssignmentEmpty(Head, 4, 5)
	PrintAllList(Head)
}
