package main

import "fmt"

// 给定一个二进制数组 nums 和一个整数 k，如果可以翻转最多 k 个 0 ，则返回 数组中连续 1 的最大个数 。
func main() {
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
}

//输入：nums = [1,1,1,0,0,0,1,1,1,1,0], K = 2
//输出：6
//解释：[1,1,1,0,0,1,1,1,1,1,1]
//粗体数字从 0 翻转到 1，最长的子数组长度为 6。

// 输入：nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
// 输出：10
// 解释：[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
// 粗体数字从 0 翻转到 1，最长的子数组长度为 10。
func longestOnes(nums []int, k int) int {
	left := 0
	m := 0
	zero := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]^0 == 0 {
			zero++
		}
		for zero > k {
			if nums[left]^0 == 0 {
				zero--
			}
			left++
		}
		if m < i-left+1 {
			m = i - left + 1
		}
	}
	return m
}
