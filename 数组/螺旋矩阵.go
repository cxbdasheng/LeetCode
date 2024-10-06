package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}))
}
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	if len(matrix) == 1 {
		return matrix[0]
	}
	left, right, up, down := 0, len(matrix[0])-1, 0, len(matrix)-1
	res := make([]int, 0)
	for {
		for i := left; i <= right; i++ {
			res = append(res, matrix[up][i])
		}
		up++
		// 判断是否越界
		if up > down {
			break
		}

		for i := up; i <= down; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		// 判断是否越界
		if right < left {
			break
		}
		// 遍历下边界
		// 从右到左加入结果集
		for i := right; i >= left; i-- {
			res = append(res, matrix[down][i])
		}
		// 下边界收缩
		down--
		// 判断是否越界
		if down < up {
			break
		}

		// 遍历左边界
		// 从下到上加入结果集
		for i := down; i >= up; i-- {
			res = append(res, matrix[i][left])
		}
		// 左边界收缩
		left++
		// 判断是否越界
		if left > right {
			break
		}
	}

	return res
}
