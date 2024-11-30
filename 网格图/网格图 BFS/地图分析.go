package main

import "fmt"

// 你现在手里有一份大小为 n x n 的 网格 grid，上面的每个 单元格 都用 0 和 1 标记好了。其中 0 代表海洋，1 代表陆地。
// 请你找出一个海洋单元格，这个海洋单元格到离它最近的陆地单元格的距离是最大的，并返回该距离。如果网格上只有陆地或者海洋，请返回 -1。
// 我们这里说的距离是「曼哈顿距离」（ Manhattan Distance）：(x0, y0) 和 (x1, y1) 这两个单元格之间的距离是 |x0 - x1| + |y0 - y1| 。
func main() {
	fmt.Println(maxDistance([][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}))
}

// 输入：grid = [[1,0,1],[0,0,0],[1,0,1]]
// 输出：2
// 解释：
// 海洋单元格 (1, 1) 和所有陆地单元格之间的距离都达到最大，最大距离为 2。

// 输入：grid = [[1,0,0],[0,0,0],[0,0,0]]
// 输出：4
// 解释：
// 海洋单元格 (2, 2) 和所有陆地单元格之间的距离都达到最大，最大距离为 4。
func maxDistance(grid [][]int) int {
	// bfs 初始化队列
	queue := make([][]int, 0)

	// 第一轮 多源入队列
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				queue = append(queue, []int{i, j})
			}
		}
	}
	if len(queue) == 0 || len(queue) == len(grid)*len(grid[0]) {
		return -1
	}
	direction := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	res := -1
	for len(queue) != 0 {
		size := len(queue)
		for i := 1; i <= size; i++ {
			point := queue[0]
			queue = queue[1:]
			for _, v := range direction {
				x := point[0] + v[0]
				y := point[1] + v[1]
				if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == 0 {
					grid[x][y] = 1
					queue = append(queue, []int{x, y})
				}
			}
		}
		res++
	}
	return res
}
