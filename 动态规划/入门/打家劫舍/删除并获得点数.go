package main

import (
	"fmt"
)

// 给你一个整数数组 nums ，你可以对它进行一些操作。
// 每次操作中，选择任意一个 nums[i] ，删除它并获得 nums[i] 的点数。之后，你必须删除 所有 等于 nums[i] - 1 和 nums[i] + 1 的元素。
// 开始你拥有 0 个点数。返回你能通过这些操作获得的最大点数。
func main() {
	fmt.Println(deleteAndEarn([]int{1, 1, 1}))
}

// 输入：nums = [3,4,2]
// 输出：6
// 解释：
// 删除 4 获得 4 个点数，因此 3 也被删除。
// 之后，删除 2 获得 2 个点数。总共获得 6 个点数。

// 输入：nums = [2,2,3,3,3,4]
// 输出：9
// 解释：
// 删除 3 获得 3 个点数，接着要删除两个 2 和 4 。
// 之后，再次删除 3 获得 3 个点数，再次删除 3 获得 3 个点数。
// 总共获得 9 个点数。
func deleteAndEarn(nums []int) int {
	var n int
	for _, num := range nums {
		n = max(n, num)
	}
	sum := make([]int, n+1)
	for _, num := range nums {
		sum[num] += num
	}
	// f[i] = max(f[i - 2] + w, f[i - 1])
	l, r := sum[0], sum[1]
	for i := 2; i <= n; i++ {
		l, r = r, max(l+sum[i], r)
	}
	return r

}
