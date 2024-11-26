package main

import "fmt"

// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
// 此外，你可以假设该网格的四条边均被水包围。
func main() {
	fmt.Println(numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}))
}

// 输入：grid = [
// ["1","1","1","1","0"],
// ["1","1","0","1","0"],
// ["1","1","0","0","0"],
// ["0","0","0","0","0"]
// ]
// 输出：1

// 输入：grid = [
// ["1","1","0","0","0"],
// ["1","1","0","0","0"],
// ["0","0","1","0","0"],
// ["0","0","0","1","1"]
// ]
// 输出：3
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	var ans int
	var dfs func(row, col int)
	dfs = func(row, col int) {
		// 出界，或者不是 '1'，就不再往下递归
		if row < 0 || row >= m || col < 0 || col >= n || grid[row][col] != '1' {
			return
		}
		grid[row][col] = '2'
		dfs(row, col-1)
		dfs(row, col+1)
		dfs(row-1, col)
		dfs(row+1, col)
	}
	for i, row := range grid {
		for j, c := range row {
			if c == '1' {
				dfs(i, j)
				ans++
			}
		}
	}
	return ans
}
