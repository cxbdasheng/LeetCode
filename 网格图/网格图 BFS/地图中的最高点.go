package main

import "fmt"

// 给你一个大小为 m x n 的整数矩阵 isWater ，它代表了一个由 陆地 和 水域 单元格组成的地图。
// 如果 isWater[i][j] == 0 ，格子 (i, j) 是一个 陆地 格子。
// 如果 isWater[i][j] == 1 ，格子 (i, j) 是一个 水域 格子。
// 你需要按照如下规则给每个单元格安排高度：
// 每个格子的高度都必须是非负的。
// 如果一个格子是 水域 ，那么它的高度必须为 0 。
// 任意相邻的格子高度差 至多 为 1 。当两个格子在正东、南、西、北方向上相互紧挨着，就称它们为相邻的格子。（也就是说它们有一条公共边）
// 找到一种安排高度的方案，使得矩阵中的最高高度值 最大 。
// 请你返回一个大小为 m x n 的整数矩阵 height ，其中 height[i][j] 是格子 (i, j) 的高度。如果有多种解法，请返回 任意一个 。
func main() {
	fmt.Println(highestPeak([][]int{{0, 1}, {0, 0}}))
}

// 输入：isWater = [[0,1],[0,0]]
// 输出：[[1,0],[2,1]]
// 解释：上图展示了给各个格子安排的高度。
// 蓝色格子是水域格，绿色格子是陆地格。

// 输入：isWater = [[0,0,1],[1,0,0],[0,0,0]]
// 输出：[[1,1,0],[0,1,1],[1,2,2]]
// 解释：所有安排方案中，最高可行高度为 2 。
// 任意安排方案中，只要最高高度为 2 且符合上述规则的，都为可行方案。
func highestPeak(isWater [][]int) [][]int {
	n := len(isWater)
	m := len(isWater[0])
	ans := make([][]int, n)
	type pair struct {
		x, y int
	}
	dir := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var queue []pair
	for i := 0; i < n; i++ {
		ans[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if isWater[i][j] == 1 {
				queue = append(queue, pair{i, j})
			}
		}
	}

	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			top := queue[0]
			queue = queue[1:]
			for _, p := range dir {
				if top.x+p.x >= 0 && top.x+p.x <= n-1 && top.y+p.y >= 0 && top.y+p.y <= m-1 && isWater[top.x+p.x][top.y+p.y] == 0 {
					isWater[top.x+p.x][top.y+p.y] = 1
					ans[top.x+p.x][top.y+p.y] = ans[top.x][top.y] + 1
					queue = append(queue, pair{top.x + p.x, top.y + p.y})
				}
			}
		}
	}
	return ans
}
