package main

import "fmt"

// 给你一个整数数组 nums 和两个正整数 m 和 k 。
//请你返回 nums 中长度为 k 的 几乎唯一 子数组的 最大和 ，如果不存在几乎唯一子数组，请你返回 0 。
//如果 nums 的一个子数组有至少 m 个互不相同的元素，我们称它是 几乎唯一 子数组。
//子数组指的是一个数组中一段连续 非空 的元素序列。

// 输入：nums = [2,6,7,3,1,7], m = 3, k = 4
// 输出：18
// 解释：总共有 3 个长度为 k = 4 的几乎唯一子数组。分别为 [2, 6, 7, 3] ，[6, 7, 3, 1] 和 [7, 3, 1, 7] 。这些子数组中，和最大的是 [2, 6, 7, 3] ，和为 18 。
func main() {
	fmt.Println(maxSum([]int{2, 6, 7, 3, 1, 7}, 3, 4))
}
func maxSum(nums []int, m int, k int) int64 {
	ans := int64(0)
	mp := make(map[int]int, k+1)
	sum := int64(0)
	for i := 0; i < k; i++ {
		mp[nums[i]]++
		sum += int64(nums[i])
	}
	if len(mp) >= m {
		ans = sum
	}
	for i := k; i < len(nums); i++ {
		sum += int64(nums[i])
		sum -= int64(nums[i-k])
		mp[nums[i]]++
		mp[nums[i-k]]--
		if mp[nums[i-k]] == 0 {
			delete(mp, nums[i-k])
		}
		if len(mp) >= m {
			if sum > ans {
				ans = sum
			}
		}
	}
	return ans
}
