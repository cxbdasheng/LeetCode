package main

import (
	"fmt"
)

func main() {
	fmt.Println(generate(5))
}
func generate(numRows int) [][]int {
	var matrix [][]int
	for i := 0; i < numRows; i++ {
		tmp := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if i < 2 {
				tmp[j] = 1
				continue
			}
			if j == 0 || j == i {
				tmp[j] = 1
				continue
			}
			if j >= 1 && j < i {
				tmp[j] = matrix[i-1][j-1] + matrix[i-1][j]
				continue
			}
		}
		matrix = append(matrix, tmp)
	}

	return matrix
}
