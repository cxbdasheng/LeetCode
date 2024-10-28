package main

import (
	"fmt"
)

// 给定一个含有 n 个正整数的数组和一个正整数 target 。
// 找出该数组中满足其总和大于等于 target 的长度最小的 子数组
// [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
func main() {
	fmt.Println(minSubArrayLen(11, []int{1, 2, 3, 4, 5}))
}

//输入：target = 7, nums = [2,3,1,2,4,3]
//输出：2
//解释：子数组 [4,3] 是该条件下的长度最小的子数组。

//输入：target = 4, nums = [1,4,4]
//输出：1

// 输入：target = 11, nums = [1,1,1,1,1,1,1,1]
// 输出：0
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	left := 0
	ans := n + 1
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		for sum >= target {
			// 总和大于等于 target 的长度最小的 子数组
			if ans > i-left+1 {
				ans = i - left + 1
			}
			sum -= nums[left]
			left++
		}
	}
	if ans <= n {
		return ans
	}
	return 0
}
