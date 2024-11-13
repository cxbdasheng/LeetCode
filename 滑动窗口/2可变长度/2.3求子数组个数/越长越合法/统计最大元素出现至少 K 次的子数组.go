package main

import "fmt"

// 给你一个整数数组 nums 和一个 正整数 k 。
//
// 请你统计有多少满足 「 nums 中的 最大 元素」至少出现 k 次的子数组，并返回满足这一条件的子数组的数目。
//
// 子数组是数组中的一个连续元素序列。
func main() {
	fmt.Println(countSubarrays([]int{1, 3, 2, 3, 3}, 2))
}

func countSubarrays(nums []int, k int) int64 {
	left := 0
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		if m < nums[i] {
			m = nums[i]
		}
	}
	var ans int64
	count := 0
	for i := 0; i < len(nums); i++ {
		if m == nums[i] {
			count++
		}
		for count >= k {
			ans += int64(len(nums)) - int64(i)
			if nums[left] == m {
				count--
			}
			left++
		}

	}
	return ans
}
