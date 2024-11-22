package main

import (
	"fmt"
)

func main() {
	var nums = []int{1, 5, 11, 5}
	fmt.Println(canPartition(nums))
	fmt.Println(canPartition1(nums))
}

func canPartition(nums []int) bool {
	var sum int
	for _, n := range nums {
		sum += n
	}
	if sum&1 == 1 {
		return false
	}
	target := sum / 2
	dp := make([]int, target+1)
	for i := 0; i < len(nums); i++ { //物品
		for j := target; j >= nums[i]; j-- {
			if dp[j] < dp[j-nums[i]]+nums[i] {
				dp[j] = dp[j-nums[i]] + +nums[i]
			}
		}
	}
	return dp[target] == target
}

func canPartition1(nums []int) bool {
	var sum int
	for _, n := range nums {
		sum += n
	}
	if sum&1 == 1 {
		return false
	}
	target := sum / 2
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if i < 0 {
			return j == 0
		}
		return j >= nums[i] && dfs(i-1, j-nums[i]) || dfs(i-1, j)
	}
	return dfs(len(nums)-1, target)
}
