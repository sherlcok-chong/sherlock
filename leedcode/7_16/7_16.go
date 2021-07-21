package __16

func search(nums []int, target int) int {
	right := len(nums)
	ans:=0
	for left := 0; left < right;  {
		if target > (left+right)/2 {
			left = (left + right) / 2
		} else if target < (left+right)/2{
			right = (left + right) / 2
		} else {
			for i := (left+right)/2; nums[i] ==target ; i++ {
				ans++
			}
			for i := (left+right)/2; nums[i] ==target ; i-- {
				ans++
			}
		}
	}
	return ans-1
}
