package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums = []int{1, 2, 1}
	fmt.Println(permuteUnique(nums))
}
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	var path []int
	use := make([]int, len(nums))
	var backtracking func([]int)
	backtracking = func(use []int) {
		if len(path) == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, path)
			result = append(result, temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i-1] == nums[i] && use[i-1] == 0 {
				continue
			}
			// 取过的就不取
			if use[i] == 0 {
				path = append(path, nums[i])
				use[i] = 1
				backtracking(use)
				path = path[:len(path)-1]
				use[i] = 0
			}
		}
	}
	backtracking(use)
	return result
}
