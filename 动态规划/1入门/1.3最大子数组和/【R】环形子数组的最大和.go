package main

import (
	"fmt"
	"math"
)

// 给定一个长度为 n 的环形整数数组 nums ，返回 nums 的非空 子数组 的最大可能和 。
// 环形数组 意味着数组的末端将会与开头相连呈环状。形式上， nums[i] 的下一个元素是 nums[(i + 1) % n] ， nums[i] 的前一个元素是 nums[(i - 1 + n) % n] 。
// 子数组 最多只能包含固定缓冲区 nums 中的每个元素一次。形式上，对于子数组 nums[i], nums[i + 1], ..., nums[j] ，不存在 i <= k1, k2 <= j 其中 k1 % n == k2 % n 。
func main() {
	fmt.Println(maxSubarraySumCircular([]int{-3, -2, -3}))
}

// 输入：nums = [1,-2,3,-2]
// 输出：3
// 解释：从子数组 [3] 得到最大和 3

// 输入：nums = [5,-3,5]
// 输出：10
// 解释：从子数组 [5,5] 得到最大和 5 + 5 = 10

// 输入：nums = [3,-2,2,-3]
// 输出：3
// 解释：从子数组 [3] 和 [3,-2,2] 都可以得到最大和 3
func maxSubarraySumCircular(nums []int) int {

	n := len(nums)
	r, l, sum := 0, 0, 0
	var minF int
	/**
	防止为负数
	*/
	maxF := math.MinInt
	for i := 0; i < n; i++ {
		r = max(r, 0) + nums[i]
		maxF = max(maxF, r)
		l = min(l, 0) + nums[i]
		minF = min(minF, l)
		sum += nums[i]
	}
	if sum == minF {
		return maxF
	}
	return max(maxF, sum-minF)
}
