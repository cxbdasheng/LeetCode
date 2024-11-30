package main

import "fmt"

// 给你一个大小为 n x n 的二元矩阵 grid ，其中 1 表示陆地，0 表示水域。
// 岛 是由四面相连的 1 形成的一个最大组，即不会与非组内的任何其他 1 相连。grid 中 恰好存在两座岛 。
// 你可以将任意数量的 0 变为 1 ，以使两座岛连接起来，变成 一座岛 。
// 返回必须翻转的 0 的最小数目。

func main() {
	fmt.Println(shortestBridge([][]int{{0, 1, 0}, {0, 0, 0}, {0, 0, 1}}))
}

// 输入：grid = [[0,1],[1,0]]
// 输出：1

// 输入：grid = [[0,1,0],[0,0,0],[0,0,1]]
// 输出：2

// 输入：grid = [[1,1,1,1,1],[1,0,0,0,1],[1,0,1,0,1],[1,0,0,0,1],[1,1,1,1,1]]
// 输出：1
func shortestBridge(grid [][]int) int {
	type pair struct {
		x, y int
	}
	dir := []pair{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	n := len(grid)
	m := len(grid[0])
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= n || y < 0 || y >= m || grid[x][y] == 0 || grid[x][y] == 2 {
			return
		}
		grid[x][y] = 2
		for _, p := range dir {
			dfs(x+p.x, y+p.y)
		}
		return
	}
	flag := false
	for i := 0; i < len(grid); i++ {
		if flag {
			break
		}
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
				flag = true
				break
			}
		}

	}
	var queue []pair
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, pair{i, j})
			}
		}
	}
	step := 0
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			top := queue[0]
			queue = queue[1:]
			for _, p := range dir {
				if top.x+p.x >= 0 && top.x+p.x <= n-1 && top.y+p.y >= 0 && top.y+p.y <= m-1 && (grid[top.x+p.x][top.y+p.y] == 1 || grid[top.x+p.x][top.y+p.y] == 0) {
					if grid[top.x+p.x][top.y+p.y] == 1 {
						return step
					} else {
						grid[top.x+p.x][top.y+p.y] = 2
					}
					queue = append(queue, pair{top.x + p.x, top.y + p.y})
				}
			}
		}
		step++
	}
	return 0
}
