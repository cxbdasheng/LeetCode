package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
func subsets(nums []int) [][]int {
	var res [][]int
	var path []int
	var backtracking func(nums []int, start int)
	backtracking = func(nums []int, start int) {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtracking(nums, i+1)
			path = path[:len(path)-1]
		}
	}
	backtracking(nums, 0)

	return res
}
