package main

import "fmt"

// 给你一个整数数组 nums 和一个整数 k ，请你返回 nums 中 好 子数组的数目。
// 一个子数组 arr 如果有 至少 k 对下标 (i, j) 满足 i < j 且 arr[i] == arr[j] ，那么称它是一个 好 子数组。
// 子数组 是原数组中一段连续 非空 的元素序列。
func main() {
	fmt.Println(countGood([]int{3, 1, 4, 3, 2, 2, 4}, 2))
}

// 输入：nums = [1,1,1,1,1], k = 10
// 输出：1
// 解释：唯一的好子数组是这个数组本身。

// 输入：nums = [3,1,4,3,2,2,4], k = 2
// 输出：4
// 解释：总共有 4 个不同的好子数组：
// - [3,1,4,3,2,2] 有 2 对。
// - [3,1,4,3,2,2,4] 有 3 对。
// - [1,4,3,2,2,4] 有 2 对。
// - [4,3,2,2,4] 有 2 对。
func countGood(nums []int, k int) int64 {
	var ans int64
	cnt := map[int]int{}
	pairs, left := 0, 0
	for _, x := range nums {
		// 重点累加 推到公式 c(x, 2) - c(x-1, 2) = x - 1
		pairs += cnt[x]
		cnt[x]++
		for pairs >= k {
			cnt[nums[left]]--
			pairs -= cnt[nums[left]]
			left++
		}
		// 左边的子数组全算
		ans += int64(left)
	}

	return ans
}
