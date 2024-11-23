package main

import (
	"fmt"
)

// 我们定义 arr 是 山形数组 当且仅当它满足：
// arr.length >= 3
// 存在某个下标 i （从 0 开始） 满足 0 < i < arr.length - 1 且：
// arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
// arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
//
// 给你整数数组 nums ，请你返回将 nums 变成 山形状数组 的 最少 删除次数。
func main() {
	fmt.Println(minimumMountainRemovals([]int{100, 92, 89, 77, 74, 66, 64, 66, 64}))
}

// 输入：nums = [1,3,1]
// 输出：0
// 解释：数组本身就是山形数组，所以我们不需要删除任何元素。

// 输入：nums = [2,1,1,5,6,2,3,1]
// 输出：3
// 解释：一种方法是将下标为 0，1 和 5 的元素删除，剩余元素为 [1,5,6,3,1] ，是山形数组。
func minimumMountainRemovals(nums []int) int {
	n := len(nums)
	dpLeft := make([]int, n)
	dpRight := make([]int, n)
	for i := 0; i < n; i++ {
		dpLeft[i] = 1
		dpRight[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dpLeft[i] = max(dpLeft[i], dpLeft[j]+1)
			}
		}

	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if nums[i] > nums[j] {
				dpRight[i] = max(dpRight[i], dpRight[j]+1)
			}
		}
	}
	var ans int
	for i := 1; i < n-1; i++ {
		if dpLeft[i] >= 2 && dpRight[i] >= 2 {
			ans = max(ans, dpLeft[i]+dpRight[i])
		}
	}
	return n - ans + 1
}
