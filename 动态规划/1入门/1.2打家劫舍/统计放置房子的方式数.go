package main

import "fmt"

const mod = 1_000_000_007

// 一条街道上共有 n * 2 个 地块 ，街道的两侧各有 n 个地块。每一边的地块都按从 1 到 n 编号。每个地块上都可以放置一所房子。
// 现要求街道同一侧不能存在两所房子相邻的情况，请你计算并返回放置房屋的方式数目。由于答案可能很大，需要对 109 + 7 取余后再返回。
// 注意，如果一所房子放置在这条街某一侧上的第 i 个地块，不影响在另一侧的第 i 个地块放置房子。
func main() {
	/**
	递推
	*/
	fmt.Println(countHousePlacements(1))
}

// 输入：n = 1
// 输出：4
// 解释：
// 可能的放置方式：
// 1. 所有地块都不放置房子。
// 2. 一所房子放在街道的某一侧。
// 3. 一所房子放在街道的另一侧。
// 4. 放置两所房子，街道两侧各放置一所。
func countHousePlacements(n int) int {
	if n == 0 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % mod
	}
	return (dp[n] * dp[n]) % mod
}
