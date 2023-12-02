package main

import "fmt"

func main() {
	var n = 10
	fmt.Println(integerBreak(n))
}
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[2] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
