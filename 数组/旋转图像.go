package main

import "fmt"

func main() {
	rotate([][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}})
}
func rotate(matrix [][]int) {
	//n := 10
	//// 创建一个 n x n 的矩阵
	//mat := make([][]int, n)
	//for i := range mat {
	//	mat[i] = make([]int, n)
	//}
	//
	//// 填充矩阵
	//for i := 0; i < n; i++ {
	//	for j := 0; j < n; j++ {
	//		mat[i][j] = i*n + j + 1 // 填充示例数据
	//	}
	//}
	//
	//// 打印矩阵
	//for _, row := range mat {
	//	fmt.Println(row)
	//}
	//matrix = mat
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] =
				matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]

			matrix[i][j] = matrix[n-j-1][i]
			matrix[n-j-1][i] = matrix[n-j-1][i]
		}
	}
	fmt.Println()
	for _, row := range matrix {
		fmt.Println(row)
	}
}
