package main

import "fmt"

// https://leetcode.cn/problems/set-matrix-zeroes/solutions/?envType=study-plan-v2&envId=top-100-liked
func main() {
	setZeroes([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}})
}
func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	mmp := make(map[int]bool, m)
	nmp := make(map[int]bool, m)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				mmp[i] = true
				nmp[j] = true
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mmp[i] || nmp[j] {
				matrix[i][j] = 0
			}
		}
	}
	fmt.Println(matrix)
}
