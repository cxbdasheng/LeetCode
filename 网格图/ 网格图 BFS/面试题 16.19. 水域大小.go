package main

import (
	"fmt"
	"sort"
)

// 你有一个用于表示一片土地的整数矩阵land，该矩阵中每个点的值代表对应地点的海拔高度。
// 若值为0则表示水域。由垂直、水平或对角连接的水域为池塘。池塘的大小是指相连接的水域的个数。
// 编写一个方法来计算矩阵中所有池塘的大小，返回值需要从小到大排序。
func main() {
	fmt.Println(pondSizes([][]int{{0, 2, 1, 0}, {0, 1, 0, 1}, {1, 1, 0, 1}, {0, 1, 0, 1}}))
}

// 输入：
// [
// [0,2,1,0],
// [0,1,0,1],
// [1,1,0,1],
// [0,1,0,1]
// ]
// 输出： [1,2,4]

func pondSizes(land [][]int) []int {
	var ans []int
	n := len(land)
	m := len(land[0])
	var dfs func(int, int) int
	dfs = func(x, y int) int {
		if x < 0 || x >= n || y < 0 || y >= m || land[x][y] > 0 {
			return 0
		}
		land[x][y] = 1
		return dfs(x+1, y) + dfs(x-1, y) + dfs(x, y+1) + dfs(x, y-1) + dfs(x-1, y-1) + dfs(x+1, y+1) + dfs(x-1, y+1) + dfs(x+1, y-1) + 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			res := dfs(i, j)
			if res > 0 {
				ans = append(ans, res)
			}
		}
	}
	sort.Ints(ans)
	return ans
}
