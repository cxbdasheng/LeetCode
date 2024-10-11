package main

import (
	"fmt"
)

var n int

const M = int64(1e9 + 7)

func main() {
	fmt.Scan(&n)
	fmt.Println(solve())
}
func solve() int64 {
	if n <= 3 {
		return int64(n)
	}
	result := power(n - 3)
	result2 := (((3 * result[0][0]) % M) + ((2 * result[1][0]) % M) + ((1 * result[2][0]) % M)) % M
	return result2
}
func power(k int) [][]int64 {
	ret := [][]int64{[]int64{1, 0, 0}, []int64{0, 1, 0}, []int64{0, 0, 1}}
	matrix := [][]int64{[]int64{1, 1, 0}, []int64{0, 0, 1}, []int64{1, 0, 0}}
	for k != 0 {
		if k&1 != 0 {
			ret = mulMatrix(ret, matrix)
		}
		matrix = mulMatrix(matrix, matrix)
		k >>= 1
	}
	return ret
}
func mulMatrix(matrix1, matrix2 [][]int64) [][]int64 {
	n, m := len(matrix1), len(matrix2[0])
	result := make([][]int64, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int64, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < len(matrix1[0]); k++ {
				result[i][j] = (result[i][j] + (matrix1[i][k]*matrix2[k][j])%M) % M
			}
		}
	}
	return result
}
