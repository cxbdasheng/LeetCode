package main

import (
	"fmt"
	"math"
	"slices"
)

// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
// 子数组 是数组中的一个连续部分。
func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray1([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

//输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
//输出：6
//解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

//输入：nums = [1]
//输出：1

// 输入：nums = [5,4,-1,7,8]
// 输出：23
/**
动态规划
*/
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1], 0) + nums[i]
	}
	return slices.Max(dp)
}

/**
前缀和
*/

func maxSubArray1(nums []int) int {
	ans := math.MinInt
	minPreSum := 0
	preSum := 0
	for _, x := range nums {
		preSum += x                        // 当前的前缀和
		ans = max(ans, preSum-minPreSum)   // 减去前缀和的最小值
		minPreSum = min(minPreSum, preSum) // 维护前缀和的最小值
	}
	return ans
}
