package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums = []int{1, 2, 1}
	sort.Ints(nums)
	use := make([]int, len(nums))
	fmt.Println(backtracking(nums, use))
}

var result [][]int
var path []int

func backtracking(nums []int, use []int) [][]int {
	if len(path) == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		result = append(result, tmp)
		return result
	}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && use[i-1] == 0 {
			continue
		}
		if use[i] == 0 {
			path = append(path, nums[i])
			use[i] = 1
			backtracking(nums, use)
			use[i] = 0
			path = path[:len(path)-1]
		}
	}
	return result
}
