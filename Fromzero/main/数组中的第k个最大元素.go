package main

import "sort"

func main() {
	matrix := make([][2]int, 20)
	//fmt.Println(matrix)
	//xor := make([][len(matrix[1])]int,len(matrix))
	m:=len(matrix)
	n:=len(matrix[1])
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			matrix[i][j] = matrix[i][j] ^ matrix[i][j-1]
		}
	}
	for i := 0; i < n; i++ {
		for j := 1; j < m; j++ {
			matrix[j][i] = matrix[j][i] ^ matrix[j-1][i]
		}
	}
	xor := make([]int,m*n)
	cnt:=0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			xor[cnt]=matrix[i][j]
			cnt++
		}
	}
	sort.Ints(xor)

}
