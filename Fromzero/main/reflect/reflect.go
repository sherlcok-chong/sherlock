package main

import (
	"fmt"
	"reflect"
)

type person struct {
	name string
	age int
	sex string
}
func (p person) fir(msg string){
	fmt.Println(msg)
}
func (p person)printInfo()  {
	fmt.Println(p.name,p.age,p.sex)
}
func GetMessage(input interface{}){
	getType := reflect.TypeOf(input)
	fmt.Println(getType.Name())
	fmt.Println(getType.Kind())

	getValue := reflect.ValueOf(input)
	fmt.Println(getValue)
	for i:= 0;i<getType.NumField();i++{
		filed:= getType.Field(i)
		//value := getValue.Field(i).Interface()
		fmt.Println(filed.Name,filed.Type )
	}
}
func main() {
	p1 := person{"qwe",12,"qwe"}
	GetMessage(p1)
}
