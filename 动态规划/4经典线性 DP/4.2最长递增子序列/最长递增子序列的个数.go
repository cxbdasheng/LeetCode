package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(findNumberOfLIS([]int{1, 3, 5, 4, 7}))
}
func findNumberOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	mp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		mp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			/**
			因为是子序列所以可以使用 dp[i] = max(dp[i], dp[j]+1)
			*/
			if nums[j] < nums[i] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
					mp[i] = mp[j]
				} else if dp[i] == dp[j]+1 {
					mp[i] += mp[j]
				}
			}
		}
	}
	var ans int
	m := slices.Max(dp)
	for i := 0; i < len(dp); i++ {
		if m == dp[i] {
			ans += mp[i]
		}

	}
	return ans
}
