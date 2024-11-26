package main

import "fmt"

// 给你一个大小为 m x n 的二进制矩阵 grid 。
// 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在 水平或者竖直的四个方向上 相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。
// 岛屿的面积是岛上值为 1 的单元格的数目。
// 计算并返回 grid 中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。
func main() {
	fmt.Println(maxAreaOfIsland([][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}))
}

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var ans int
	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		// 出界，或者不是 '1'，就不再往下递归
		if row < 0 || row >= m || col < 0 || col >= n {
			return 0
		}

		if grid[row][col] != 1 {
			return 0
		}

		grid[row][col] = 2

		return 1 + dfs(row+1, col) + dfs(row-1, col) + dfs(row, col-1) + dfs(row, col+1)
	}
	for i, row := range grid {
		for j, col := range row {
			if col == 1 {
				ans = max(ans, dfs(i, j))
			}
		}
	}
	return ans
}
