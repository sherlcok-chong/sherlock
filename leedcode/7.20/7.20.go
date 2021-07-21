package __20

import (
	"sort"
)

func minPairSum(nums []int) int {
	sort.Ints(nums)
	ans:=9999999
	for i,r := 0,len(nums); i < r; i++ {
		sum:=nums[i]+nums[r]
		if ans>sum {
			ans=sum
		}
	}
	return ans
}
