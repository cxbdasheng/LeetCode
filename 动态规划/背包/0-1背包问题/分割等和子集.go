package main

import "fmt"

func main() {
	var nums = []int{1, 5, 11, 5}
	fmt.Println(canPartition(nums))
}

func canPartition(nums []int) bool {
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	dp := make([]int, target+1)
	dp[0] = 0
	// i 为物品 j为背包重量
	for i := 0; i < len(nums); i++ { //物品
		for j := target; j >= nums[i]; j-- { //背包
			if dp[j] < dp[j-nums[i]]+nums[i] {
				dp[j] = dp[j-nums[i]] + +nums[i]
			}
		}
	}
	return dp[target] == target
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
