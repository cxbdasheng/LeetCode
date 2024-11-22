package main

import (
	"fmt"
	"slices"
)

// 给你一个整数数组 rewardValues，长度为 n，代表奖励的值。
// 最初，你的总奖励 x 为 0，所有下标都是 未标记 的。你可以执行以下操作 任意次 ：
// 从区间 [0, n - 1] 中选择一个 未标记 的下标 i。
// 如果 rewardValues[i] 大于 你当前的总奖励 x，则将 rewardValues[i] 加到 x 上（即 x = x + rewardValues[i]），并 标记 下标 i。
// 以整数形式返回执行最优操作能够获得的 最大 总奖励。
func main() {
	fmt.Println(maxTotalReward([]int{1, 1, 3, 3}))
}
func maxTotalReward(rewardValues []int) int {
	slices.Sort(rewardValues)
	n := len(rewardValues)
	mx := rewardValues[n-1] * 2 // 最大总奖励 = rewardValues[n-1]*2-1
	dp := make([][]bool, n+1)   // 初始化 dp 数组
	for i := range dp {
		dp[i] = make([]bool, mx)
	}
	dp[0][0] = true
	// dp[i][j], i 代表选到第 i 个元素，j 代表目前的总奖励数
	for i := 1; i <= n; i++ {
		x := rewardValues[i-1]
		for j := 0; j < mx; j++ {
			dp[i][j] = dp[i-1][j]    // 不选当前元素，直接转移，什么都不改变
			if j-x >= 0 && x > j-x { // 只有 x 大于当前总奖励才能选
				// 假设 dp[i-1][j] 是 true，那可以直接返回 true，无所谓选不选
				dp[i][j] = dp[i-1][j] || dp[i-1][j-x] // 选当前元素
			}
		}
	}
	ans := mx - 1             // 最好的情况就是 mx-1
	for dp[n][ans] == false { // 找到 == true 就是答案
		ans--
	}
	return ans
}
