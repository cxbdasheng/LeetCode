package main

import "fmt"

// 给你一个大小为 m x n 的矩阵 grid 。最初，你位于左上角 (0, 0) ，每一步，你可以在矩阵中 向右 或 向下 移动。
// 在从左上角 (0, 0) 开始到右下角 (m - 1, n - 1) 结束的所有路径中，找出具有 最大非负积 的路径。路径的积是沿路径访问的单元格中所有整数的乘积。
// 返回 最大非负积 对 109 + 7 取余 的结果。如果最大积为 负数 ，则返回 -1 。
// 注意，取余是在得到最大积之后执行的。
func main() {
	fmt.Println(maxProductPath([][]int{{4, 3}}))
}

// 输入：grid = [[-1,-2,-3],[-2,-3,-3],[-3,-3,-2]]
// 输出：-1
// 解释：从 (0, 0) 到 (2, 2) 的路径中无法得到非负积，所以返回 -1 。

// 输入：grid = [[1,-2,1],[1,-2,1],[3,-4,1]]
// 输出：8
// 解释：最大非负积对应的路径如图所示 (1 * 1 * -2 * -4 * 1 = 8)

// 输入：grid = [[1,3],[0,-4]]
// 输出：0
// 解释：最大非负积对应的路径如图所示 (1 * 0 * -4 = 0)
func maxProductPath(grid [][]int) int {
	const mod = 1_000_000_007
	maxDp := make([][]int, len(grid))
	minDp := make([][]int, len(grid))
	for i := range maxDp {
		maxDp[i] = make([]int, len(grid[i]))
		minDp[i] = make([]int, len(grid[i]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j == 0 {
				maxDp[i][j] = grid[i][j]
				minDp[i][j] = grid[i][j]
				continue
			}
			if i == 0 {
				maxDp[i][j] = maxDp[i][j-1] * grid[i][j]
				minDp[i][j] = minDp[i][j-1] * grid[i][j]
				continue
			}
			if j == 0 {
				maxDp[i][j] = maxDp[i-1][j] * grid[i][j]
				minDp[i][j] = minDp[i-1][j] * grid[i][j]
				continue
			}
			if grid[i][j] >= 0 {
				maxDp[i][j] = max(maxDp[i-1][j], maxDp[i][j-1]) * grid[i][j]
				minDp[i][j] = min(minDp[i-1][j], minDp[i][j-1]) * grid[i][j]
			} else {
				maxDp[i][j] = min(minDp[i-1][j], minDp[i][j-1]) * grid[i][j]
				minDp[i][j] = max(maxDp[i-1][j], maxDp[i][j-1]) * grid[i][j]
			}

		}
	}
	ans := max(maxDp[len(maxDp)-1][len(maxDp[0])-1], minDp[len(minDp)-1][len(minDp[0])-1])
	if ans < 0 {
		return -1
	}
	return ans % mod
}
