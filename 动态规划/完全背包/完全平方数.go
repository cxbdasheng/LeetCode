package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numSquares(12))
}

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for j := 1; j <= n; j++ {
		dp[j] = math.MaxUint32
	}
	// 物品
	for i := 1; i*i <= n; i++ {
		for j := i * i; j <= n; j++ {
			// 如果是 MaxUint32 则证明 [j-i*i] 没有匹配到
			if dp[j-i*i] != math.MaxUint32 {
				// 推导公式
				dp[j] = Min(dp[j], dp[j-i*i]+1)
			}
		}
	}
	return dp[n]
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
