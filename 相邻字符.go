package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)
	dp := make([]int, n)
	maps := make(map[int]bool)
	dp[0] = 0
	for i := 1; i < n; i++ {
		if s[i] == s[i-1] && !maps[i-1] {
			dp[i] = dp[i-1] + 1
			maps[i] = true
			continue
		}

		dp[i] = dp[i-1]
	}
	fmt.Print(dp[n-1])
}
