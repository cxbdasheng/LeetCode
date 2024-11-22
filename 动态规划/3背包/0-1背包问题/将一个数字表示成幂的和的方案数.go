package main

import (
	"fmt"
)

// 给你两个 正 整数 n 和 x 。
// 请你返回将 n 表示成一些 互不相同 正整数的 x 次幂之和的方案数。换句话说，你需要返回互不相同整数 [n1, n2, ..., nk] 的集合数目，满足 n = n1x + n2x + ... + nkx 。
// 由于答案可能非常大，请你将它对 109 + 7 取余后返回。
// 比方说，n = 160 且 x = 3 ，一个表示 n 的方法是 n = 23 + 33 + 53 。
func main() {
	fmt.Println(numberOfWays(4, 1))
	fmt.Println(numberOfWays1(4, 1))
}

// 输入：n = 10, x = 2
// 输出：1
// 解释：我们可以将 n 表示为：n = 32 + 12 = 10 。
// 这是唯一将 10 表达成不同整数 2 次方之和的方案。

// 输入：n = 4, x = 1
// 输出：2
// 解释：我们可以将 n 按以下方案表示：
// - n = 41 = 4 。
// - n = 31 + 11 = 4 。
func numberOfWays(n int, x int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; pow(i, x) <= n; i++ {
		v := pow(i, x)
		for s := n; s >= v; s-- {
			dp[s] += dp[s-v]
		}
	}
	return dp[n] % 1_000_000_007
}
func pow(x, y int) int {
	if y == 0 {
		return 1
	}
	ans := 1
	for y > 0 {
		if y&1 == 1 {
			ans *= x
		}
		x *= x
		y >>= 1
	}
	return ans
}

/**
DFS
*/

func numberOfWays1(n int, x int) int {
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if j == 0 {
			return 1
		}
		if pow(i, x) > j {
			return 0
		}
		res := dfs(i+1, j-pow(i, x)) % 1_000_000_007
		res = (res + dfs(i+1, j)) % 1_000_000_007

		return res
	}
	return dfs(1, n)
}
