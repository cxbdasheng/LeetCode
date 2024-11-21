package main

import "fmt"

// 给你一个下标从 0 开始、大小为 m x n 的矩阵 grid ，矩阵由若干 正 整数组成。
// 你可以从矩阵第一列中的 任一 单元格出发，按以下方式遍历 grid ：
// 从单元格 (row, col) 可以移动到 (row - 1, col + 1)、(row, col + 1) 和 (row + 1, col + 1) 三个单元格中任一满足值 严格 大于当前单元格的单元格。
// 返回你在矩阵中能够 移动 的 最大 次数。
func main() {
	fmt.Println(maxMoves([][]int{{3, 2, 4}, {2, 1, 9}, {1, 1, 7}}))
}

// 输入：grid = [[2,4,3,5],[5,4,9,3],[3,4,2,11],[10,9,13,15]]
// 输出：3
// 解释：可以从单元格 (0, 0) 开始并且按下面的路径移动：
// - (0, 0) -> (0, 1).
// - (0, 1) -> (1, 2).
// - (1, 2) -> (2, 3).
// 可以证明这是能够移动的最大次数。

// 输入：grid = [[3,2,4],[2,1,9],[1,1,7]]
// 输出：0
// 解释：从第一列的任一单元格开始都无法移动。
func maxMoves(grid [][]int) int {
	ans := 1
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[i]))
	}
	// 列
	for i := 0; i < len(grid[0]); i++ {
		// 行
		for j := 0; j < len(grid); j++ {
			if i == 0 {
				dp[j][i] = 1
				continue
			}
			if j-1 >= 0 && j+1 < len(grid) {
				if (grid[j][i] > grid[j-1][i-1] && dp[j-1][i-1] > 0) || (grid[j][i] > grid[j][i-1] && dp[j][i-1] > 0) || (grid[j][i] > grid[j+1][i-1] && dp[j+1][i-1] > 0) {
					dp[j][i] = i + 1
				}
			} else if j == 0 && j+1 < len(grid) {
				if (grid[j][i] > grid[j+1][i-1] && dp[j+1][i-1] > 0) || (grid[j][i] > grid[j][i-1] && dp[j][i-1] > 0) {
					dp[j][i] = i + 1
				}
			} else if j == len(grid)-1 && j-1 >= 0 {
				if (grid[j][i] > grid[j][i-1] && dp[j][i-1] > 0) || (grid[j][i] > grid[j-1][i-1] && dp[j-1][i-1] > 0) {
					dp[j][i] = i + 1
				}
			} else {
				if grid[j][i] > grid[j][i-1] && grid[j][i-1] > 0 {
					dp[j][i] = i + 1
				}
			}
			if ans < dp[j][i] {
				ans = dp[j][i]
			}
		}
	}
	return ans - 1
}
