package main

import (
	"fmt"
	"math"
)

// 给你一个下标从 0 开始的整数数组 nums 和一个整数 target 。
// 返回和为 target 的 nums 子序列中，子序列 长度的最大值 。如果不存在和为 target 的子序列，返回 -1 。
// 子序列 指的是从原数组中删除一些或者不删除任何元素后，剩余元素保持原来的顺序构成的数组。
func main() {
	fmt.Println(lengthOfLongestSubsequence([]int{2, 3, 5}, 5))
	fmt.Println(lengthOfLongestSubsequence1([]int{2, 3, 5}, 5))
}

// 输入：nums = [1,2,3,4,5], target = 9
// 输出：3
// 解释：总共有 3 个子序列的和为 9 ：[4,5] ，[1,3,5] 和 [2,3,4] 。最长的子序列是 [1,3,5] 和 [2,3,4] 。所以答案为 3 。

// 输入：nums = [4,1,3,2,1,5], target = 7
// 输出：4
// 解释：总共有 5 个子序列的和为 7 ：[4,3] ，[4,1,2] ，[4,2,1] ，[1,1,5] 和 [1,3,2,1] 。最长子序列为 [1,3,2,1] 。所以答案为 4 。

// 输入：nums = [1,1,5,4,5], target = 3
// 输出：-1
// 解释：无法得到和为 3 的子序列。

// 题解：https://leetcode.cn/problems/length-of-the-longest-subsequence-that-sums-to-target/solutions/2502839/mo-ban-qia-hao-zhuang-man-xing-0-1-bei-b-0nca/
func lengthOfLongestSubsequence(nums []int, target int) int {
	n := len(nums)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, target+1)
		for j := range dp[i] {
			dp[i][j] = math.MinInt
		}
	}
	dp[0][0] = 0
	for i, num := range nums {
		for c := 0; c <= target; c++ {
			if c < num {
				dp[i+1][c] = dp[i][c]
			} else {
				dp[i+1][c] = max(dp[i][c], dp[i][c-num]+1)
			}
		}
	}
	if dp[n][target] > 0 {
		return dp[n][target]
	}
	return -1
}

func lengthOfLongestSubsequence1(nums []int, target int) int {
	n := len(nums)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, target+1)
		for j := 0; j <= target; j++ {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, c int) int {
		if i < 0 {
			if c == 0 {
				return 0
			}
			return math.MinInt
		}
		if memo[i][c] != -1 {
			return memo[i][c]
		}
		if c < nums[i] {
			memo[i][c] = dfs(i-1, c)
		} else {
			memo[i][c] = max(dfs(i-1, c), dfs(i-1, c-nums[i])+1)
		}
		return memo[i][c]
	}
	ans := dfs(n-1, target)
	if ans > 0 {
		return ans
	}
	return -1
}
