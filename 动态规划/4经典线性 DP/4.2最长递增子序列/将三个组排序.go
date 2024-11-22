package main

import (
	"fmt"
	"slices"
)

// 给你一个整数数组 nums 。nums 的每个元素是 1，2 或 3。在每次操作中，你可以删除 nums 中的一个元素。返回使 nums 成为 非递减 顺序所需操作数的 最小值。
func main() {
	fmt.Println(minimumOperations([]int{2, 1, 3, 2, 1}))
}

// 输入：nums = [2,1,3,2,1]
// 输出：3
// 解释：
// 其中一个最优方案是删除 nums[0]，nums[2] 和 nums[3]。

// 输入：nums = [1,3,2,1,3,3]
// 输出：2
// 解释：
// 其中一个最优方案是删除 nums[1] 和 nums[2]。

// 输入：nums = [2,2,2,2,3,3]
// 输出：0
// 解释：
// nums 已是非递减顺序的。
func minimumOperations(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if i != j && nums[j] <= nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return len(dp) - slices.Max(dp)
}
