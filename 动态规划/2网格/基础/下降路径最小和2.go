package main

import (
	"fmt"
	"math"
	"slices"
)

// 给你一个 n x n 整数矩阵 grid ，请你返回 非零偏移下降路径 数字和的最小值。
// 非零偏移下降路径 定义为：从 grid 数组中的每一行选择一个数字，且按顺序选出来的数字中，相邻数字不在原数组的同一列。
func main() {
	fmt.Println(minFallingPathSum([][]int{{2, 2, 1, 2, 2}, {2, 2, 1, 2, 2}, {2, 2, 1, 2, 2}, {2, 2, 1, 2, 2}, {2, 2, 1, 2, 2}}))
}

// 输入：grid = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：13
// 解释：
// 所有非零偏移下降路径包括：
// [1,5,9], [1,5,7], [1,6,7], [1,6,8],
// [2,4,8], [2,4,9], [2,6,7], [2,6,8],
// [3,4,8], [3,4,9], [3,5,7], [3,5,9]
// 下降路径中数字和最小的是 [1,5,7] ，所以答案是 13 。

// 输入：grid = [[7]]
// 输出：7
func minFallingPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[i]))

	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 {
				dp[i][j] = grid[i][j]
				continue
			}
			dp[i][j] = math.MaxInt
			for k := 0; k < len(grid[i]); k++ {
				if k == j {
					continue
				}
				dp[i][j] = min(dp[i][j], dp[i-1][k])
			}
			dp[i][j] = dp[i][j] + grid[i][j]
		}
	}
	return slices.Min(dp[len(dp)-1])
}
