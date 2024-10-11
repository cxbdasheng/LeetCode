package main

import "fmt"

func main() {
	fmt.Println(searchMatrix([][]int{{1, 3}}, 3))

}
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	x, y := 0, len(matrix[0])-1
	for x < m && y >= 0 {
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] > target {
			y--
		} else {
			x++
		}
	}

	return false
}
