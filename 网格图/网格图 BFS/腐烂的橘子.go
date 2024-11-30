package main

import "fmt"

// 在给定的 m x n 网格 grid 中，每个单元格可以有以下三个值之一：
// 值 0 代表空单元格；
// 值 1 代表新鲜橘子；
// 值 2 代表腐烂的橘子。
// 每分钟，腐烂的橘子 周围 4 个方向上相邻 的新鲜橘子都会腐烂。
// 返回 直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1 。
func main() {
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}))
}

// 输入：grid = [[2,1,1],[1,1,0],[0,1,1]]
// 输出：4

// 输入：grid = [[2,1,1],[0,1,1],[1,0,1]]
// 输出：-1
// 解释：左下角的橘子（第 2 行， 第 0 列）永远不会腐烂，因为腐烂只会发生在 4 个方向上。
//
// 输入：grid = [[0,2]]
// 输出：0
// 解释：因为 0 分钟时已经没有新鲜橘子了，所以答案就是 0 。
func orangesRotting(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	type pair struct {
		x, y int
	}
	dir := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var queue []pair
	var count int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				count++
			}
			if grid[i][j] == 2 {
				queue = append(queue, pair{i, j})
			}
		}
	}
	ans := 0
	for len(queue) > 0 {
		l := len(queue)
		flag := false
		for i := 0; i < l; i++ {
			top := queue[0]
			queue = queue[1:]
			for _, p := range dir {
				if top.x+p.x >= 0 && top.x+p.x <= n-1 && top.y+p.y >= 0 && top.y+p.y <= m-1 && grid[top.x+p.x][top.y+p.y] == 1 {
					grid[top.x+p.x][top.y+p.y]++
					count--
					flag = true
					queue = append(queue, pair{top.x + p.x, top.y + p.y})
				}
			}
		}
		if flag {
			ans++
		}
	}
	if count > 0 {
		return -1
	}
	return ans
}
