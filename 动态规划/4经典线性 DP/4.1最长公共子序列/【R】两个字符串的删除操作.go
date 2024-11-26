package main

import "fmt"

// 给定两个单词 word1 和 word2 ，返回使得 word1 和  word2 相同所需的最小步数。
// 每步 可以删除任意一个字符串中的一个字符。
func main() {
	fmt.Println(minDistance("sea", "eat"))
	fmt.Println(minDistance1("sea", "eat"))
}

// 输入: word1 = "sea", word2 = "eat"
// 输出: 2
// 解释: 第一步将 "sea" 变为 "ea" ，第二步将 "eat "变为 "ea"

// 输入：word1 = "leetcode", word2 = "etco"
// 输出：4

/*
*
减去最长公共子序列
*/
func minDistance(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return n + m - dp[n][m]*2
}

/*
*
动态规划直接做
*/
func minDistance1(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word2)+1)
	}
	//初始化
	for i := 0; i < len(dp); i++ {
		dp[i][0] = i
	}
	for j := 0; j < len(dp[0]); j++ {
		dp[0][j] = j
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			// 两个字符相等的情况下，结果等于两个串的前一个字符相加
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// dp[i-1][j]+1 为删除 word1[i-1]
				// dp[i][j-1]+1 为删除 word2[j-1]
				// dp[i-1][j-1]+2（两个都删除） 为删除 word1[i-1] 和 word2[j-1]
				dp[i][j] = min(min(dp[i-1][j]+1, dp[i][j-1]+1), dp[i-1][j-1]+2)
				// dp[i][j-1]= min(dp[i-1][j-1]+1,dp[i][j-2]+1,dp[i-1][j-2]+2)
				// 也就是说如果是  dp[i-1][j-1]+2 最小的话
				// dp【i-1】【j-1】 + 2这个状态是有可能从 dp【i-i】【j】+1 这层的状态里的组合出来的，所以+2这个是重复了
				// dp【i-1】【j-1】 + 2 == dp【i-i】【j】+1， 要么dp【i-1】【j-1】 + 2 > dp【i-i】【j】+1, 所以放在min里时， dp【i-1】【j-1】 + 2 怎么都不会被选到。

			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}
