package main

import "fmt"

func main() {
	fmt.Println(searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5))
}
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	i, j := 0, n-1        // 从右上角开始
	for i < m && j >= 0 { // 还有剩余元素
		if matrix[i][j] == target {
			return true // 找到 target
		}
		if matrix[i][j] < target {
			i++ // 这一行剩余元素全部小于 target，排除
		} else {
			j-- // 这一列剩余元素全部大于 target，排除
		}
	}
	return false
}
