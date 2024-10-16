package main

import "fmt"

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j >= 1 {
				dp[i][j] = grid[i][j] + dp[i][j-1]
			}
			if j == 0 && i >= 1 {
				dp[i][j] = grid[i][j] + dp[i-1][j]
			}
			if i >= 1 && j >= 1 {
				dp[i][j] = min(dp[i-1][j]+grid[i][j], dp[i][j-1]+grid[i][j])
			}
		}
	}
	return dp[m-1][n-1]
}
