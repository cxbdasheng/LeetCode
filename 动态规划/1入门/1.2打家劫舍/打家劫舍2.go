package main

import "fmt"

// 你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。
func main() {
	num := []int{1, 2, 3, 1}
	result1 := rob(num[1:])
	result2 := rob(num[:len(num)-1])
	fmt.Println(max(result1, result2))
}

// 输入：nums = [2,3,2]
// 输出：3
// 解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的

// 输入：nums = [1,2,3,1]
// 输出：4
// 解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
// 偷窃到的最高金额 = 1 + 3 = 4 。

// 输入：nums = [1,2,3]
// 输出：3
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(dp[0], nums[1])
	for i := 2; i < len(dp); i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}
	return max(dp[len(nums)-1], dp[len(nums)-2])
}
