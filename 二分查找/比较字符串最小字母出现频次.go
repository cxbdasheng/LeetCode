package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numSmallerByFrequency([]string{"bba", "abaaaaaa", "aaaaaa", "bbabbabaab", "aba", "aa", "baab", "bbbbbb", "aab", "bbabbaabb"}, []string{"aaabbb", "aab", "babbab", "babbbb", "b", "bbbbbbbbab", "a", "bbbbbbbbbb", "baaabbaab", "aa"}))
}

// calFrequency 辅助函数：计算单词中最小字符出现的频率 f(s)。
// 这是对原 calTransfer 函数的修正和重命名。
func calFrequency(s string) int {
	if len(s) == 0 {
		return 0
	}
	minChar := byte('z' + 1)
	count := 0
	for i := 0; i < len(s); i++ {
		char := s[i]
		if minChar > char {
			minChar = char
			count = 1
		} else if minChar == char {
			count++
		}
	}
	return count
}

func numSmallerByFrequency(queries []string, words []string) []int {
	ans := make([]int, len(queries))
	nums := make([]int, len(words))
	for i := 0; i < len(words); i++ {
		nums[i] = calFrequency(words[i])
	}
	sort.Ints(nums)
	for i := 0; i < len(queries); i++ {
		ans[i] = len(words) - sort.SearchInts(nums, calFrequency(queries[i])+1)
	}
	return ans
}
