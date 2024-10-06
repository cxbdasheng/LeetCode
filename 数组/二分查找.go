package main

import "fmt"

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
}

// 时间复杂度 O(logn)
func search(nums []int, target int) int {
	// 初始化左右边界
	left := 0
	right := len(nums)

	// 循环逐步缩小区间范围
	for left < right {
		// 求区间中点
		mid := left + (right-left)>>1

		// 根据 nums[mid] 和 target 的大小关系
		// 调整区间范围
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	// 在输入数组内没有找到值等于 target 的元素
	return -1
}
