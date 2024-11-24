package main

import "fmt"

// 给两个整数数组 nums1 和 nums2 ，返回 两个数组中 公共的 、长度最长的子数组的长度 。
func main() {
	fmt.Println(findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}))
}

// 输入：nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
// 输出：3
// 解释：长度最长的公共子数组是 [3,2,1] 。

// 输入：nums1 = [0,0,0,0,0], nums2 = [0,0,0,0,0]
// 输出：5
func findLength(nums1 []int, nums2 []int) int {
	n1 := len(nums1)
	n2 := len(nums2)
	dp := make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
	}
	var ans int
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if ans < dp[i][j] {
					ans = dp[i][j]
				}
			}
		}
	}
	return ans
}
