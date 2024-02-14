package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{3, 2, 6, 5, 0, 3}))
}

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 4)
	}

	// 买入状态：
	// 操作一：前一天就是持有股票状态（状态一），dp[i][0] = dp[i - 1][0]
	// 操作二：今天买入了，有两种情况
	// 前一天是冷冻期（状态四），dp[i - 1][3] - prices[i]
	// 前一天是保持卖出股票的状态（状态二），dp[i - 1][1] - prices[i]
	dp[0][0] = -prices[0]
	// 保持卖出股票状态:1.前一天就是状态二;2.前一天是冷冻期（状态四）
	dp[0][1] = 0
	// 今天卖出状态:前一天是冷冻期
	dp[0][2] = 0
	// 冷冻期：前一天卖出状态
	dp[0][3] = 0

	for i := 1; i < n; i++ {
		dp[i][0] = Max(dp[i-1][0], Max(dp[i-1][3]-prices[i], dp[i-1][1]-prices[i]))
		dp[i][1] = Max(dp[i-1][1], dp[i-1][3])
		dp[i][2] = dp[i-1][0] + prices[i]
		dp[i][3] = dp[i-1][2]
	}
	return Max(dp[n-1][3], max(dp[n-1][1], dp[n-1][2]))
}
func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
