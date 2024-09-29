package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{1, 2, 1}))
}
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i-1] + nums[i]
		}
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}
