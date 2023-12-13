package main

import "fmt"

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if inDict(s[j:i], wordDict) && dp[j] == true {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}

func inDict(s string, wordDict []string) bool {
	for i := 0; i < len(wordDict); i++ {
		if s == wordDict[i] {
			return true
		}
	}
	return false
}
