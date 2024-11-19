package main

import "fmt"

func main() {
	num := []int{1, 2, 3, 1}
	result1 := rob(num[1:])
	result2 := rob(num[:len(num)-1])
	fmt.Println(Max(result1, result2))
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = Max(dp[0], nums[1])
	for i := 2; i < len(dp); i++ {
		dp[i] = Max(dp[i-1], nums[i]+dp[i-2])
	}
	return Max(dp[len(nums)-1], dp[len(nums)-2])
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
