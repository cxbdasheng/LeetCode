package main

import (
	"fmt"
	"sort"
)

// 元素的 频数 是该元素在一个数组中出现的次数。
//
// 给你一个整数数组 nums 和一个整数 k 。在一步操作中，你可以选择 nums 的一个下标，并将该下标对应元素的值增加 1 。
//
// 执行最多 k 次操作后，返回数组中最高频元素的 最大可能频数 。
func main() {
	fmt.Println(maxFrequency([]int{1, 2, 4}, 5))
}

// 输入：nums = [1,2,4], k = 5
// 输出：3
// 解释：对第一个元素执行 3 次递增操作，对第二个元素执 2 次递增操作，此时 nums = [4,4,4] 。
// 4 是数组中最高频元素，频数是 3 。

//输入：nums = [1,4,8,13], k = 5
//输出：2
//解释：存在多种最优解决方案：
//- 对第一个元素执行 3 次递增操作，此时 nums = [4,4,8,13] 。4 是数组中最高频元素，频数是 2 。
//- 对第二个元素执行 4 次递增操作，此时 nums = [1,8,8,13] 。8 是数组中最高频元素，频数是 2 。
//- 对第三个元素执行 5 次递增操作，此时 nums = [1,4,13,13] 。13 是数组中最高频元素，频数是 2 。

// 入：nums = [3,9,6], k = 2
// 输出：1

// 题解：https://leetcode.cn/problems/frequency-of-the-most-frequent-element/solutions/742905/jian-dan-yi-dong-zui-zi-ran-de-chu-li-lu-9i9a/
func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	left := 0
	ans := 1
	sum := 0
	for i := 1; i < len(nums); i++ {
		// 当前区间总频率数 = 前面一个数的差值 * 当前的区间大小
		// 例子：1,2,4
		// 【1-4 区间的频率数】 = （4-2）*（3-0） + 【1-3 区间的频率数】
		sum += (nums[i] - nums[i-1]) * (i - left)
		for sum > k {
			sum -= nums[i] - nums[left]
			left++
		}
		ans = max(ans, i-left+1)
	}

	return ans
}
