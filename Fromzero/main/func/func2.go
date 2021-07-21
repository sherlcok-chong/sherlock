package main

import "fmt"

func main() {
	func() {
		fmt.Println("给句痛快话")
	}()
	f1()
	f2 := f1
	f2()
	func(a, b int) {
		fmt.Println("投降不投降", a, b)
	}(1, 2)
	i:=func(a, b int) int {
		return a + b
	}(1,2)
	fmt.Println(i)
}
func f1() {
	fmt.Println("山本，我日你仙人")
}
