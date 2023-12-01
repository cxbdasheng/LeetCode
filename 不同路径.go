package main

import "fmt"

func main() {
	var m, n = 3, 7
	fmt.Println(uniquePaths(m, n))
}
func uniquePaths(m int, n int) int {

	if m == 0 || n == 0 {
		return 0
	}
	if n == 1 || m == 1 {
		return 1
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = 0
				continue
			}
			if i == 0 {
				dp[i][j] = 1
				continue
			}
			if j == 0 {
				dp[i][j] = 1
				continue
			}
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}
	return dp[m-1][n-1]
}
