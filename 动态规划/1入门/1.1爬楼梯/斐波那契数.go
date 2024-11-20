package main

import "fmt"

func main() {
	var n = 42421631561
	fmt.Println(fib(n))
}
func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}
