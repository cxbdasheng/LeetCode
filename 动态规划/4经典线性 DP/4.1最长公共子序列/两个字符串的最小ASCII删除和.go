package main

import "fmt"

func main() {
	fmt.Println(minimumDeleteSum("delete", "leet"))
}
func minimumDeleteSum(s1 string, s2 string) int {
	n1 := len(s1)
	n2 := len(s2)
	dp := make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
	}
	for i := 1; i <= n1; i++ {
		dp[i][0] = int(s1[i-1]) + dp[i-1][0]
	}
	for i := 1; i <= n2; i++ {
		dp[0][i] = int(s2[i-1]) + dp[0][i-1]
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+int(s1[i-1]), dp[i][j-1]+int(s2[j-1]))
			}
		}
	}
	return dp[n1][n2]
}
