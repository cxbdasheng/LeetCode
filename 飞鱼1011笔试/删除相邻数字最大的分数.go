package main

import "fmt"

func main() {
	fmt.Println(boredom([]int{1, 2, 1, 3, 2, 2, 2, 2, 3}))
}
func boredom(a []int) int {
	// write code here
	m := max(a...)
	dp := make([]int, m+1)
	for _, v := range a {
		dp[v] += v
	}
	for i := 2; i <= m; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+dp[i])
	}
	fmt.Println(dp)
	return dp[m]
}

func max(nums ...int) int {
	m := nums[0]
	for _, v := range nums {
		if v > m {
			m = v
		}
	}
	return m
}
