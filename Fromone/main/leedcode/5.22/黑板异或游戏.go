package main

func main() {

}
func xorGame(nums []int) bool {
	var sum int
	for i,_:= range nums{
		sum^=i
	}
	return sum==0||len(nums)%2==0
}