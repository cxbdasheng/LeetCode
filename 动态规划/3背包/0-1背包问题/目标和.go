package main

import (
	"fmt"
	"math"
)

// 将nums中元素分为左右两个部分，一部分是正集合left，另一部分是负集合right
// 例子：【1,1,1,1,1】
// 由题意可知 1.正数集合left - 负数集合right = 目标值 traget  比如：【1,1,1,1】 - 【1】 = 3
// 2.正数集合left + 负数集合right = 数组总和 sum  比如： 【1,1,1,1】 + 【1】 = 5
// 又因为 我们只要得知凑出left正集合有多少种组合方式 就等同于题目问题：凑成目标和有多少种方式
// 例如 【-1,1,1,1,1】 【1,-1,1,1,1】 【1,1,-1,1,1】 【1,1,1,-1,1】 【1,1,1,1,-1】
// 于是相当于说每个元素放不放进正数的集合里 等同于背包问题中的拿或不拿
// 所以我们要的结果是：数组中能够凑成背包容量为正集合left的最大组合方案数
// 而left可以用可知的1和2代入解出为 left = (sum + target) / 2
// 然而left代表的是正数集要多少个才能凑成目标和target(正数集个数确定负数集也就确认) 除以2存在不能整除情况
func main() {
	var nums = []int{1, 1, 1, 1, 1}
	fmt.Println(findTargetSumWays(nums, 3))
}
func findTargetSumWays(nums []int, target int) int {
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if (sum+target)%2 > 0 || int(math.Abs(float64(target))) > sum {
		return 0
	}
	left := (sum + target) / 2
	dp := make([]int, left+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := left; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[left]
}
