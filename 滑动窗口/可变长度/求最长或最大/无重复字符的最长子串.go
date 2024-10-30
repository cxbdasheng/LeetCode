package main

import "fmt"

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。
func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

// 输入: s = "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
// 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	mp := make(map[byte]int)
	left := 0
	m := 0
	for i := 0; i < len(s); i++ {
		//当前进入窗口的大于 1 则遍历
		for mp[s[i]] >= 1 {
			mp[s[left]]--
			left++
		}
		mp[s[i]]++
		if m < i-left+1 {
			m = i - left + 1
		}
	}
	return m
}
