package main

func countTriplets(arr []int) int {
	ans :=0
	len := len(arr)
	xor := make([]int,len+1)
	for i,val :=range arr{
		xor[i+1]=xor[i]^val
	}
	for i:=0;i<len;i++ {
		for j:=i+1;j<len;j++{
			for k:=j;k<len;k++ {
				if xor[i]==xor[k+1]{
					ans++
				}
			}
		}
	}
	return ans
}
