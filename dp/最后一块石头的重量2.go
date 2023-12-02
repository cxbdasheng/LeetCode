package main

import "fmt"

func main() {
	var nums = []int{2, 7, 4, 1, 8, 1}
	fmt.Println(lastStoneWeightII(nums))
}

func lastStoneWeightII(stones []int) int {
	var sum int
	for i := 0; i < len(stones); i++ {
		sum += stones[i]
	}
	target := sum / 2
	dp := make([]int, target+1)
	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			if dp[j] < dp[j-stones[i]]+stones[i] {
				dp[j] = dp[j-stones[i]] + stones[i]
			}
		}
	}
	return sum - dp[target] - dp[target]
}
