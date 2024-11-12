package main

import (
	"fmt"
	"math"
)

// 给你一个下标从 0 开始的数组 nums 和一个整数 target 。
// 下标从 0 开始的数组 infinite_nums 是通过无限地将 nums 的元素追加到自己之后生成的。
// 请你从 infinite_nums 中找出满足 元素和 等于 target 的 最短 子数组，并返回该子数组的长度。如果不存在满足条件的子数组，返回 -1 。
func main() {
	fmt.Println(minSizeSubarray([]int{1, 2, 2, 2, 1, 2, 1, 2, 1, 2, 1}, 83))
}

// 输入：nums = [1,2,3], target = 5
// 输出：2
// 解释：在这个例子中 infinite_nums = [1,2,3,1,2,3,1,2,...] 。
// 区间 [1,2] 内的子数组的元素和等于 target = 5 ，且长度 length = 2 。
// 可以证明，当元素和等于目标值 target = 5 时，2 是子数组的最短长度。

// 输入：nums = [1,1,1,2,3], target = 4
// 输出：2
// 解释：在这个例子中 infinite_nums = [1,1,1,2,3,1,1,1,2,3,1,1,...].
// 区间 [4,5] 内的子数组的元素和等于 target = 4 ，且长度 length = 2 。
// 可以证明，当元素和等于目标值 target = 4 时，2 是子数组的最短长度。

// 输入：nums = [2,4,6,8], target = 3
// 输出：-1
// 解释：在这个例子中 infinite_nums = [2,4,6,8,2,4,6,8,...] 。
// 可以证明，不存在元素和等于目标值 target = 3 的子数组。
func minSizeSubarray(nums []int, target int) int {
	total := 0
	for _, x := range nums {
		total += x
	}
	ans := math.MaxInt
	left, sum, n := 0, 0, len(nums)
	for right := 0; right < n*2; right++ {
		sum += nums[right%n]
		for sum > target%total {
			sum -= nums[left%n]
			left++
		}
		if sum == target%total {
			ans = min(ans, right-left+1)
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans + target/total*n
}
