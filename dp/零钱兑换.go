package main

import (
	"fmt"
	"math"
)

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount))
}
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for j := 1; j <= amount; j++ {
		dp[j] = math.MaxUint32
	}
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxUint32 {
				// 推导公式
				dp[j] = Min(dp[j], dp[j-coins[i]]+1)
			}
			//if dp[j] != 0 {
			//	dp[j] = Min(dp[j], dp[j-coins[i]]+1)
			//} else {
			//	dp[j] = dp[j-coins[i]] + 1
			//}
		}
	}
	if dp[amount] == math.MaxUint32 {
		return -1
	}
	return dp[amount]
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
