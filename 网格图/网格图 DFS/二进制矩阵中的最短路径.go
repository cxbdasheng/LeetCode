package main

import "fmt"

// 给你一个 n x n 的二进制矩阵 grid 中，返回矩阵中最短 畅通路径 的长度。如果不存在这样的路径，返回 -1 。
// 二进制矩阵中的 畅通路径 是一条从 左上角 单元格（即，(0, 0)）到 右下角 单元格（即，(n - 1, n - 1)）的路径，该路径同时满足下述要求：
// 路径途经的所有单元格的值都是 0 。
// 路径中所有相邻的单元格应当在 8 个方向之一 上连通（即，相邻两单元之间彼此不同且共享一条边或者一个角）。
// 畅通路径的长度 是该路径途经的单元格总数。
func main() {
	fmt.Println(shortestPathBinaryMatrix([][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}}))
}

// 输入：grid = [[0,1],[1,0]]
// 输出：2

// 输入：grid = [[0,0,0],[1,1,0],[1,1,0]]
// 输出：4

// 输入：grid = [[1,0,0],[1,1,0],[1,1,0]]
// 输出：-1
func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] != 0 {
		return -1
	}
	n := len(grid)
	m := len(grid[0])
	if n == 1 {
		return 1
	}
	type pair struct {
		x, y int
	}
	dir := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	var queue []pair
	queue = append(queue, pair{0, 0})

	grid[0][0] = 1
	for i := 1; len(queue) > 0; i++ {
		temp := queue
		queue = nil
		for _, p := range temp {
			for j := 0; j < 8; j++ {
				if p.x+dir[j][0] >= 0 && p.x+dir[j][0] < m && p.y+dir[j][1] >= 0 && p.y+dir[j][1] < n && grid[p.x+dir[j][0]][p.y+dir[j][1]] == 0 {
					if p.x+dir[j][0] == m-1 && p.y+dir[j][1] == n-1 {
						return i + 1
					}
					grid[p.x+dir[j][0]][p.y+dir[j][1]] = 1
					queue = append(queue, pair{p.x + dir[j][0], p.y + dir[j][1]})
				}
			}
		}
	}
	return -1
}
