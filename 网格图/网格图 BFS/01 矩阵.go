package main

import "fmt"

// 给定一个由 0 和 1 组成的矩阵 mat ，请输出一个大小相同的矩阵，其中每一个格子是 mat 中对应位置元素到最近的 0 的距离。
// 两个相邻元素间的距离为 1 。
func main() {
	fmt.Println(updateMatrix([][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}))
}

// 输入：mat = [[0,0,0],[0,1,0],[0,0,0]]
// 输出：[[0,0,0],[0,1,0],[0,0,0]]

// 输入：mat = [[0,0,0],[0,1,0],[1,1,1]]
// 输出：[[0,0,0],[0,1,0],[1,2,1]]
func updateMatrix(mat [][]int) [][]int {
	type pair struct {
		x, y int
	}
	dir := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	n := len(mat)
	m := len(mat[0])
	var queue []pair
	ans := make([][]int, n)
	for i, row := range mat {
		ans[i] = make([]int, m)
		for j, v := range row {
			if v == 0 {
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
				if top.x+p.x >= 0 && top.x+p.x <= n-1 && top.y+p.y >= 0 && top.y+p.y <= m-1 && mat[top.x+p.x][top.y+p.y] == 1 {
					mat[top.x+p.x][top.y+p.y] = 0
					ans[top.x+p.x][top.y+p.y] = ans[top.x][top.y] + 1
					queue = append(queue, pair{top.x + p.x, top.y + p.y})
				}
			}
		}
	}
	return ans
}
