package main

import "fmt"

// 给你一个二进制数组 nums ，你需要从中删掉一个元素。
// 请你在删掉元素的结果数组中，返回最长的且只包含 1 的非空子数组的长度。
// 如果不存在这样的子数组，请返回 0 。
// 输入：nums = [1,1,0,1]
// 输出：3
// 解释：删掉位置 2 的数后，[1,1,1] 包含 3 个 1 。
func main() {
	fmt.Println(longestSubarray([]int{1, 1, 1}))
}

//输入：nums = [1,1,0,1]
//输出：3
//解释：删掉位置 2 的数后，[1,1,1] 包含 3 个 1 。

//输入：nums = [0,1,1,1,0,1,1,0,1]
//输出：5
//解释：删掉位置 4 的数字后，[0,1,1,1,1,1,0,1] 的最长全 1 子数组为 [1,1,1,1,1] 。

// 输入：nums = [1,1,1]
// 输出：2
// 解释：你必须要删除一个元素。
func longestSubarray(nums []int) int {
	left := 0
	zero := 0
	ans := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zero++
		}
		for zero >= 2 {
			if nums[left] == 0 {
				zero--
			}
			left++
		}
		if i-left+1 > ans {
			ans = i - left + 1
		}
	}
	// 必须要删除一个
	if zero == 0 {
		zero++
	}
	return ans - zero
}
