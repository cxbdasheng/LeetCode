package main

import "fmt"

func main() {
	coins := []int{1, 2, 5}
	amount := 5
	fmt.Println(change(amount, coins))
}

func change(amount int, coins []int) int {
	if len(coins) <= 1 {
		return 1
	}
	var dp = make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := 0; j <= amount; j++ {
			if coins[i] <= j {
				dp[j] = dp[j] + dp[j-coins[i]]
			}
		}
	}
	return dp[amount]
}
