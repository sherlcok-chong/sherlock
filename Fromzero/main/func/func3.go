package main

import "fmt"

func main() {
	res:=added(4,5,add)
	fmt.Printf("%d\n",res)
	res2:=added(3,4,func(a,b int)int{
		return a+b
	})
	fmt.Println(res2)
}

func add(a,b int)int  {
	return a+b
}

func added(a,b int,fun func(int ,int)int)int  {
	res:= fun(a,b)
	return a+b+res
}