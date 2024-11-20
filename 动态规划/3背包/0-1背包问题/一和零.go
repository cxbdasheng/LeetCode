package main

import "fmt"

func main() {
	var strs = []string{"10", "0001", "111001", "1", "0"}
	fmt.Println(findMaxForm(strs, 5, 3))
}
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 0
	for i := 0; i < len(strs); i++ {
		x, y := 0, 0
		for j := 0; j < len(strs[i]); j++ {
			if strs[i][j] == '0' {
				x++
			} else {
				y++
			}
		}
		for k := m; k >= x; k-- {
			for l := n; l >= y; l-- {
				dp[k][l] = Max(dp[k][l], dp[k-x][l-y]+1)
			}
		}
	}
	return dp[m][n]
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
