package main

import (
	"fmt"
	"math/big"
	"slices"
)

// 给你一个整数数组 rewardValues，长度为 n，代表奖励的值。
// 最初，你的总奖励 x 为 0，所有下标都是 未标记 的。你可以执行以下操作 任意次 ：
// 从区间 [0, n - 1] 中选择一个 未标记 的下标 i。
// 如果 rewardValues[i] 大于 你当前的总奖励 x，则将 rewardValues[i] 加到 x 上（即 x = x + rewardValues[i]），并 标记 下标 i。
// 以整数形式返回执行最优操作能够获得的 最大 总奖励。
func main() {
	fmt.Println(maxTotalReward([]int{1, 1, 3, 3}))
}
func maxTotalReward(rewardValues []int) int {
	m := slices.Max(rewardValues)
	has := map[int]bool{}
	for _, v := range rewardValues {
		if v == m-1 {
			return m*2 - 1
		}
		if has[v] {
			continue
		}
		if has[m-1-v] {
			return m*2 - 1
		}
		has[v] = true
	}

	slices.Sort(rewardValues)
	rewardValues = slices.Compact(rewardValues) // 去重

	one := big.NewInt(1)
	f := big.NewInt(1)
	p := new(big.Int)
	for _, v := range rewardValues {
		mask := p.Sub(p.Lsh(one, uint(v)), one)
		f.Or(f, p.Lsh(p.And(f, mask), uint(v)))
	}
	return f.BitLen() - 1
}
