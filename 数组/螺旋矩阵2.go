package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateMatrix(3))
}

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	left, right, top, bottom := 0, n-1, 0, n-1
	var count int
	for {
		for i := left; i <= right; i++ {
			res[top][i] = count + 1
			count++
		}
		top++
		if top > bottom {
			break
		}
		for i := top; i <= bottom; i++ {
			res[i][right] = count + 1
			count++
		}
		right--
		if right < left {
			break
		}

		for i := right; i >= left; i-- {
			res[bottom][i] = count + 1
			count++
		}
		bottom--
		if bottom < top {
			break
		}

		for i := bottom; i >= top; i-- {
			res[i][left] = count + 1
			count++
		}
		left++
		if left > right {
			break
		}
	}
	return res
}
