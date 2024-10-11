package main

import "fmt"

// 红蓝染色法
func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
}
func searchRange(nums []int, target int) []int {
	start := lowerBound(nums, target) // 使用其中一种写法即可
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1} // nums 中没有 target
	}
	// 如果 start 存在，那么 end 必定存在
	end := lowerBound(nums, target+1) - 1
	return []int{start, end}
}

func lowerBound(nums []int, target int) int {
	left, right := 0, len(nums)
	var mid int
	for left < right {
		mid = left + (right-left)>>1
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
