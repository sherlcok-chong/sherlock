package main

import "sort"

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	cnt:=0
	for _,val:=range costs{
		coins-=val
		if coins<0 {
			break
		}
		cnt++
	}
	return cnt
}