package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))

}
func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 5)
	}
	// 无操作的状态
	dp[0][0] = 0
	// 第一次持有状态
	dp[0][1] = -prices[0]
	// 第一次卖出状态
	dp[0][2] = 0
	// 第二次持有状态
	dp[0][3] = dp[0][2] - prices[0]
	// 第二次卖出状态
	dp[0][4] = 0

	// max()
	for i := 1; i < len(prices); i++ {
		// 第一次持有状态
		dp[i][1] = Max(dp[i-1][0]-prices[i], dp[i-1][1])
		// 第一次卖出状态
		dp[i][2] = Max(dp[i-1][1]+prices[i], dp[i-1][2])
		// 第二次持有状态
		dp[i][3] = Max(dp[i-1][2]-prices[i], dp[i-1][3])
		// 第二次卖出状态
		dp[i][4] = Max(dp[i-1][3]+prices[i], dp[i-1][4])
	}
	return dp[len(prices)-1][4]
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
