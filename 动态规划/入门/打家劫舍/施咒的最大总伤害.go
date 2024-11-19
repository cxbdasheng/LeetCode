package main

import (
	"fmt"
	"sort"
)

// 一个魔法师有许多不同的咒语。
// 给你一个数组 power ，其中每个元素表示一个咒语的伤害值，可能会有多个咒语有相同的伤害值。
// 已知魔法师使用伤害值为 power[i] 的咒语时，他们就 不能 使用伤害为 power[i] - 2 ，power[i] - 1 ，power[i] + 1 或者 power[i] + 2 的咒语。
// 每个咒语最多只能被使用 一次 。
// 请你返回这个魔法师可以达到的伤害值之和的 最大值 。
func main() {
	fmt.Println(maximumTotalDamage([]int{5, 9, 2, 10, 2, 7, 10, 9, 3, 8}))
	fmt.Println(maximumTotalDamage1([]int{5, 9, 2, 10, 2, 7, 10, 9, 3, 8}))
}

// 输入：power = [1,1,3,4]
// 输出：6
// 解释：
// 可以使用咒语 0，1，3，伤害值分别为 1，1，4，总伤害值为 6 。

// 输入：power = [7,1,6,6]
// 输出：13
// 解释：
// 可以使用咒语 1，2，3，伤害值分别为 1，6，6，总伤害值为 13 。

/*
*
记忆化搜索
*/
func maximumTotalDamage(power []int) int64 {
	mp := map[int]int{}
	var arr []int
	for _, v := range power {
		if _, ok := mp[v]; !ok {
			arr = append(arr, v)
		}
		mp[v]++
	}
	if len(arr) == 1 {
		return int64(arr[0] * mp[arr[0]])
	}
	sort.Ints(arr)
	cache := make(map[int]int64)
	var dfs func(pos int) int64
	dfs = func(pos int) int64 {
		if pos < 0 {
			return 0
		}
		if v, ok := cache[pos]; ok {
			return v
		}
		x := arr[pos]
		j := pos
		// 必须要 j-1 否则当 pos == 0 时会无限循环
		for j > 0 && arr[j-1] >= x-2 {
			j -= 1
		}
		res := max(dfs(pos-1), dfs(j-1)+int64(x*mp[x]))
		cache[pos] = res
		return res
	}
	return dfs(len(arr) - 1)
}

/*
*
递推
*/
func maximumTotalDamage1(power []int) int64 {
	mp := map[int]int{}
	var arr []int
	for _, v := range power {
		if _, ok := mp[v]; !ok {
			arr = append(arr, v)
		}
		mp[v]++
	}
	if len(arr) == 1 {
		return int64(arr[0] * mp[arr[0]])
	}
	if len(arr) == 1 {
		return int64(arr[0] * mp[arr[0]])
	}
	dp := make([]int64, len(arr)+1)
	sort.Ints(arr)
	for i, v := range arr {
		j := i
		// 必须要 j-1 否则当 pos == 0 时会无限循环
		for j > 0 && arr[j-1] >= v-2 {
			j -= 1
		}
		dp[i+1] = max(dp[i], dp[j]+int64(v*mp[v]))
	}
	return dp[len(arr)]
}
