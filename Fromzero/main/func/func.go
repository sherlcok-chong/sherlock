package main

import "fmt"

//func getSum(num int) (n int) {
//	sum:=0
//	for i := 0; i <= num; i++ {
//		sum+=i
//	}
//	return sum
//}
func change(arr []int)(int,int){

	return arr[0],arr[1]
}
func main() {
	//n:=0
	var a = [4]int{1,2,3,4}
	arr:=a[0:3]

	//fmt.Scanf("%d",&n)
	//fmt.Println(change(1,2,3,4,5))

	//var arr = []int{1,2,3,4}
	c,d:=change(arr)
	fmt.Println(c,d,"\n",a)
}
