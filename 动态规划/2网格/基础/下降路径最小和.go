package main

import (
	"fmt"
	"math"
	"slices"
)

// 给你一个 n x n 的 方形 整数数组 matrix ，请你找出并返回通过 matrix 的下降路径 的 最小和 。
// 下降路径 可以从第一行中的任何元素开始，并从每一行中选择一个元素。在下一行选择的元素和当前行所选元素最多相隔一列（即位于正下方或者沿对角线向左或者向右的第一个元素）。
// 具体来说，位置 (row, col) 的下一个元素应当是 (row + 1, col - 1)、(row + 1, col) 或者 (row + 1, col + 1) 。
func main() {
	fmt.Println(minFallingPathSum([][]int{{17, 82}, {1, -44}}))
	fmt.Println(minFallingPathSum1([][]int{{17, 82}, {1, -44}}))
}

// 输入：matrix = [[2,1,3],[6,5,4],[7,8,9]]
// 输出：13
// 解释：如图所示，为和最小的两条下降路径

// 输入：matrix = [[-19,57],[-40,-5]]
// 输出：-59
// 解释：如图所示，为和最小的下降路径
func minFallingPathSum(matrix [][]int) int {
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[i]))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == 0 {
				dp[i][j] = matrix[i][j]
				continue
			}
			if j-1 >= 0 && j+1 < len(matrix[i-1]) {
				dp[i][j] = min(dp[i-1][j], min(dp[i-1][j+1], dp[i-1][j-1])) + matrix[i][j]
			} else if j == 0 && j+1 < len(matrix[i-1]) {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j+1]) + matrix[i][j]
			} else if j == len(matrix[i-1])-1 && j-1 >= 0 {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + matrix[i][j]
			} else {
				dp[i][j] = dp[i-1][j] + matrix[i][j]
			}
		}
	}
	return slices.Min(dp[len(dp)-1])
}

/*
*
DFS
*/
func minFallingPathSum1(matrix [][]int) int {
	n := len(matrix)
	var dfs func(int, int) int
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = math.MinInt
		}
	}
	dfs = func(i, j int) int {
		if j < 0 || j >= n { // 出界
			return math.MaxInt
		}
		if i == 0 {
			return matrix[0][j]
		}
		if memo[i][j] != math.MinInt {
			return memo[i][j]
		}
		memo[i][j] = min(min(dfs(i-1, j-1), dfs(i-1, j)), dfs(i-1, j+1)) + matrix[i][j]
		return memo[i][j]
	}
	ans := math.MaxInt
	for c := 0; c < n; c++ {
		ans = min(ans, dfs(n-1, c))
	}
	return ans
}
