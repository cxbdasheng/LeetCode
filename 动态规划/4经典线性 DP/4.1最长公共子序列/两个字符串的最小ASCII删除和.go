package main

import "fmt"

// 给定两个字符串s1 和 s2，返回 使两个字符串相等所需删除字符的 ASCII 值的最小和 。
func main() {
	fmt.Println(minimumDeleteSum("delete", "leet"))
}

// 输入: s1 = "sea", s2 = "eat"
// 输出: 231
// 解释: 在 "sea" 中删除 "s" 并将 "s" 的值(115)加入总和。
// 在 "eat" 中删除 "t" 并将 116 加入总和。
// 结束时，两个字符串相等，115 + 116 = 231 就是符合条件的最小和。

// 输入: s1 = "delete", s2 = "leet"
// 输出: 403
// 解释: 在 "delete" 中删除 "dee" 字符串变成 "let"，
// 将 100[d]+101[e]+101[e] 加入总和。在 "leet" 中删除 "e" 将 101[e] 加入总和。
// 结束时，两个字符串都等于 "let"，结果即为 100+101+101+101 = 403 。
// 如果改为将两个字符串转换为 "lee" 或 "eet"，我们会得到 433 或 417 的结果，比答案更大。
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
