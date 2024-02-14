package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	// max()
	for i := 1; i < len(prices); i++ {
		dp[i][0] = Max(dp[i-1][0], -prices[i])
		dp[i][1] = Max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[len(prices)-1][1]
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
