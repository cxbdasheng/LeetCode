package main

import (
	"fmt"
	"sort"
)

// 你打算构建一些障碍赛跑路线。给你一个 下标从 0 开始 的整数数组 obstacles ，数组长度为 n ，其中 obstacles[i] 表示第 i 个障碍的高度。
// 对于每个介于 0 和 n - 1 之间（包含 0 和 n - 1）的下标  i ，在满足下述条件的前提下，请你找出 obstacles 能构成的最长障碍路线的长度：
// 你可以选择下标介于 0 到 i 之间（包含 0 和 i）的任意个障碍。
// 在这条路线中，必须包含第 i 个障碍。
// 你必须按障碍在 obstacles 中的 出现顺序 布置这些障碍。
// 除第一个障碍外，路线中每个障碍的高度都必须和前一个障碍 相同 或者 更高 。
// 返回长度为 n 的答案数组 ans ，其中 ans[i] 是上面所述的下标 i 对应的最长障碍赛跑路线的长度。
func main() {
	fmt.Println(longestObstacleCourseAtEachPosition([]int{3, 1, 5, 6, 4, 2}))
	fmt.Println(longestObstacleCourseAtEachPosition1([]int{3, 1, 5, 6, 4, 2}))
}

// 输入：obstacles = [1,2,3,2]
// 输出：[1,2,3,3]
// 解释：每个位置的最长有效障碍路线是：
// - i = 0: [1], [1] 长度为 1
// - i = 1: [1,2], [1,2] 长度为 2
// - i = 2: [1,2,3], [1,2,3] 长度为 3
// - i = 3: [1,2,3,2], [1,2,2] 长度为 3

// 输入：obstacles = [2,2,1]
// 输出：[1,2,1]
// 解释：每个位置的最长有效障碍路线是：
// - i = 0: [2], [2] 长度为 1
// - i = 1: [2,2], [2,2] 长度为 2
// - i = 2: [2,2,1], [1] 长度为 1

// 输入：obstacles = [3,1,5,6,4,2]
// 输出：[1,1,2,3,2,2]
// 解释：每个位置的最长有效障碍路线是：
// - i = 0: [3], [3] 长度为 1
// - i = 1: [3,1], [1] 长度为 1
// - i = 2: [3,1,5], [3,5] 长度为 2, [1,5] 也是有效的障碍赛跑路线
// - i = 3: [3,1,5,6], [3,5,6] 长度为 3, [1,5,6] 也是有效的障碍赛跑路线
// - i = 4: [3,1,5,6,4], [3,4] 长度为 2, [1,4] 也是有效的障碍赛跑路线
// - i = 5: [3,1,5,6,4,2], [1,2] 长度为 2
/*
*
贪心+二分不超时
*/
func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	var ans, dp []int
	for _, v := range obstacles {
		p := sort.SearchInts(dp, v+1)
		ans = append(ans, p+1)
		if p < len(dp) {
			dp[p] = v
		} else {
			dp = append(dp, v)
		}
	}
	return ans
}

/*
*
双循环超时
*/
func longestObstacleCourseAtEachPosition1(obstacles []int) []int {
	n := len(obstacles)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			/**
			因为是子序列所以可以使用 dp[i] = max(dp[i], dp[j]+1)
			*/
			if obstacles[j] <= obstacles[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return dp
}
