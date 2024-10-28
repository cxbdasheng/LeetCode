package main

import "fmt"

// 给你一个字符串 s ，它只包含三种字符 a, b 和 c 。
// 请你返回 a，b 和 c 都 至少 出现过一次的子字符串数目。
func main() {
	fmt.Println(numberOfSubstrings("abcabc"))
}

// 输入：s = "abcabc"
// 输出：10
// 解释：包含 a，b 和 c 各至少一次的子字符串为 "abc", "abca", "abcab", "abcabc", "bca", "bcab", "bcabc", "cab", "cabc" 和 "abc" (相同字符串算多次)。

// 输入：s = "aaacb"
// 输出：3
// 解释：包含 a，b 和 c 各至少一次的子字符串为 "aaacb", "aacb" 和 "acb" 。
// 题解：https://leetcode.cn/problems/number-of-substrings-containing-all-three-characters/solutions/107665/si-kao-de-guo-cheng-bi-da-an-zhong-yao-xiang-xi-tu/
func numberOfSubstrings(s string) int {
	ans := 0
	left := 0
	mp := make([]int, 3)
	for i := 0; i < len(s); i++ {
		mp[GetIndex(s[i])]++
		for mp[0] >= 1 && mp[1] >= 1 && mp[2] >= 1 {
			ans += len(s) - i
			mp[GetIndex(s[left])]--
			left++
		}

	}
	return ans
}
func GetIndex(s byte) int {
	ans := 0
	if s == 'b' {
		ans = 1
	} else if s == 'c' {
		ans = 2
	}
	return ans
}
