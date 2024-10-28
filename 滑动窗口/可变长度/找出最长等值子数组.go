package main

import "fmt"

// 给你一个下标从 0 开始的整数数组 nums 和一个整数 k 。
// 如果子数组中所有元素都相等，则认为子数组是一个 等值子数组 。注意，空数组是 等值子数组 。
// 从 nums 中删除最多 k 个元素后，返回可能的最长等值子数组的长度。
//
// 子数组 是数组中一个连续且可能为空的元素序列。
func main() {
	fmt.Println(longestEqualSubarray([]int{1, 1, 2, 2, 1, 1}, 3))
}

// 输入：nums = [1,3,2,3,1,3], k = 3
// 输出：3
// 解释：最优的方案是删除下标 2 和下标 4 的元素。
// 删除后，nums 等于 [1, 3, 3, 3] 。
// 最长等值子数组从 i = 1 开始到 j = 3 结束，长度等于 3 。
// 可以证明无法创建更长的等值子数组。

// 输入：nums = [1,1,2,2,1,1], k = 2
// 输出：4
// 解释：最优的方案是删除下标 2 和下标 3 的元素。
// 删除后，nums 等于 [1, 1, 1, 1] 。
// 数组自身就是等值子数组，长度等于 4 。
// 可以证明无法创建更长的等值子数组。
func longestEqualSubarray(nums []int, k int) int {
	ans := 0
	posLists := make([][]int, len(nums)+1)
	for i, x := range nums {
		// 记录与上一个相同元素相隔多远
		// i 是当前元素 len(posLists[x]) 是之前总共有几个元素，i-len(posLists[x]) 相当于等于
		posLists[x] = append(posLists[x], i-len(posLists[x]))
	}
	for _, pos := range posLists {
		// 求最长，当前所有元素小于 ans 则无法让 ans 变得更大
		if len(pos) <= ans {
			continue
		}
		left := 0
		for right, p := range pos {
			// 当前元素减去上一个元素的距离等于要 删除 的元素
			// 如果要删除的元素超过 k 则++
			for p-pos[left] > k { // 要删除的数太多了
				left++
			}
			ans = max(ans, right-left+1)
		}
	}
	return ans
}
