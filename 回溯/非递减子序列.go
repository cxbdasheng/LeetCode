package main

import (
	"fmt"
)

func main() {
	fmt.Println(findSubsequences([]int{4, 6, 7, 7}))
}
func findSubsequences(nums []int) [][]int {
	var ans [][]int
	var path []int
	var backtrack func(int)
	backtrack = func(start int) {
		if len(path) > 1 {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
		}
		used := make(map[int]bool, len(nums)) // 初始化used字典，用以对同层元素去重
		for i := start; i < len(nums); i++ {
			if used[nums[i]] { // 去重
				continue
			}
			// 剪枝
			if len(path) == 0 || nums[i] >= path[len(path)-1] {
				path = append(path, nums[i])
				used[nums[i]] = true
				backtrack(i + 1)
				path = path[:len(path)-1]
			}

		}
	}
	backtrack(0)
	return ans
}
