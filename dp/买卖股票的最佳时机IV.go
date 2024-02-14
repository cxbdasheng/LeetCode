package main

import "fmt"

func main() {
	fmt.Println(maxProfit(2, []int{3, 2, 6, 5, 0, 3}))
}

func maxProfit(k int, prices []int) int {
	if k == 0 || len(prices) == 0 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2*k+1)
	}
	for i := 1; i < 2*k; i += 2 {
		dp[0][i] = -prices[0]
	}
	for i := 1; i < len(prices); i++ {
		for j := 0; j < 2*k-1; j += 2 {
			dp[i][j+1] = Max(dp[i-1][j+1], dp[i-1][j]-prices[i])
			dp[i][j+2] = Max(dp[i-1][j+2], dp[i-1][j+1]+prices[i])
		}
	}
	return dp[len(prices)-1][2*k]
}
func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
