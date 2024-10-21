package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}
func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	var path []int
	res = append(res, path)
	var backTrack func([]int, int, []int)
	backTrack = func(nums []int, start int, used []int) {
		if len(path) > 0 {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
		}
		for i := start; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && used[i-1] == 0 {
				continue
			}
			used[i] = 1
			path = append(path, nums[i])
			backTrack(nums, i+1, used)
			used[i] = 0
			path = path[:len(path)-1]
		}
	}
	used := make([]int, len(nums))
	backTrack(nums, 0, used)
	return res
}
