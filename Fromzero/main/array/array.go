package main

import "fmt"

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
	/*for i, value := range arr {
		fmt.Println(i, value)
	}*/
	var slice = []int{0, 0, 0}
	fmt.Println(slice[0])
	fmt.Println(slice)
	s3 := make([]int, 0, 10)
	fmt.Println(s3)
	fmt.Println(len(s3), cap(s3))
	s3 = append(s3, 1, 2, 3, 4, 5, 6, 7, 6, 7)
	s3 = append(s3, slice...)
	fmt.Println(s3)
	fmt.Println(len(s3), cap(s3))
	s := arr[3:5]

	fmt.Println(s)
	s[1] = 0
	fmt.Println(arr, s)
	s = append(s, slice...)
	s[3] = 0
	s[0] = 5
	fmt.Println(arr, s)
	copy(slice[0:2], s3[4:5])
	fmt.Println(slice)

}
