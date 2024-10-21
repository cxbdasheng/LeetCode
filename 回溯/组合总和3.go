package main

import "fmt"

func main() {
	fmt.Println(combinationSum3(3, 7))
}
func combinationSum3(k int, n int) [][]int {
	var ans [][]int
	var path []int
	var backtrack func(int, int, int)
	backtrack = func(k int, n int, start int) {
		if n == 0 && len(path) == k {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		for i := start; i <= 9; i++ {
			path = append(path, i)
			backtrack(k, n-i, i+1)
			path = path[:len(path)-1]
		}
	}
	backtrack(k, n, 1)
	return ans
}
