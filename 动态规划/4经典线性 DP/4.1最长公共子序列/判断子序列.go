package main

import "fmt"

// 给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
// 字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
}

// 输入：s = "abc", t = "ahbgdc"
// 输出：true

// 输入：s = "axc", t = "ahbgdc"
// 输出：false
func isSubsequence(s string, t string) bool {
	n1 := len(s)
	n2 := len(t)
	dp := make([][]int, n2+1)
	for i := 0; i <= n2; i++ {
		dp[i] = make([]int, n1+1)
	}
	for i := 1; i <= n2; i++ {
		for j := 1; j <= n1; j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n2][n1] == n1
}
