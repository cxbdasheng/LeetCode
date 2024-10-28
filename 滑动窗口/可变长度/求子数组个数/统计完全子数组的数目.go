package main

import "fmt"

// 给你一个由 正 整数组成的数组 nums 。
// 如果数组中的某个子数组满足下述条件，则称之为 完全子数组 ：
// 子数组中 不同 元素的数目等于整个数组不同元素的数目。
// 返回数组中 完全子数组 的数目。
// 子数组 是数组中的一个连续非空序列。
func main() {
	fmt.Println(countCompleteSubarrays([]int{1, 3, 1, 2, 2}))
}

//输入：nums = [1,3,1,2,2]
//输出：4
//解释：完全子数组有：[1,3,1,2]、[1,3,1,2,2]、[3,1,2] 和 [3,1,2,2] 。

// 输入：nums = [5,5,5,5]
// 输出：10
// 解释：数组仅由整数 5 组成，所以任意子数组都满足完全子数组的条件。子数组的总数为 10 。
func countCompleteSubarrays(nums []int) int {
	mp := make(map[int]int)
	ans := 0
	for i := 0; i < len(nums); i++ {
		mp[nums[i]]++
	}
	n := len(mp)
	subMp := make(map[int]int, n)
	left := 0
	for _, v := range nums { // 枚举子数组右端点 v=nums[i]
		subMp[v]++
		for len(subMp) == n {
			x := nums[left]
			subMp[x]--
			if subMp[x] == 0 {
				delete(subMp, x)
			}
			left++
		}
		ans += left // 子数组左端点 < left 的都是合法的
	}
	return ans
}
