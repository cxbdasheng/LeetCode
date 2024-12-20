package main

import (
	"fmt"
)

// 给你一个下标从 0 开始的整数数组 nums 。nums 的一个子数组如果满足以下条件，那么它是 不间断 的：
// i，i + 1 ，...，j  表示子数组中的下标。对于所有满足 i <= i1, i2 <= j 的下标对，都有 0 <= |nums[i1] - nums[i2]| <= 2 。
// 请你返回 不间断 子数组的总数目。
// 子数组是一个数组中一段连续 非空 的元素序列。

func main() {
	fmt.Println(continuousSubarrays([]int{5, 4, 2, 4}))
}

// 输入：nums = [5,4,2,4]
// 输出：8
// 解释：
// 大小为 1 的不间断子数组：[5], [4], [2], [4] 。
// 大小为 2 的不间断子数组：[5,4], [4,2], [2,4] 。
// 大小为 3 的不间断子数组：[4,2,4] 。
// 没有大小为 4 的不间断子数组。
// 不间断子数组的总数目为 4 + 3 + 1 = 8 。
// 除了这些以外，没有别的不间断子数组。

// 输入：nums = [1,2,3]
// 输出：6
// 解释：
// 大小为 1 的不间断子数组：[1], [2], [3] 。
// 大小为 2 的不间断子数组：[1,2], [2,3] 。
// 大小为 3 的不间断子数组：[1,2,3] 。
// 不间断子数组的总数目为 3 + 2 + 1 = 6 。
func continuousSubarrays(nums []int) int64 {
	var ans int64
	l := 0
	n := len(nums)
	mp := make(map[int]int)
	var maxValue, minValue int
	for i := 0; i < n; i++ {
		mp[nums[i]]++
		for {
			maxValue, minValue = nums[i], nums[i]
			for k := range mp {
				maxValue = max(maxValue, k)
				minValue = min(minValue, k)
			}
			if maxValue-minValue <= 2 {
				break
			}
			mp[nums[l]]--
			if mp[nums[l]] == 0 {
				delete(mp, nums[l])
			}
			l++
		}
		ans += int64(i - l + 1)
	}
	return ans
}
