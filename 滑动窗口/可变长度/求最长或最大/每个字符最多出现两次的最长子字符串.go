package main

import "fmt"

// 给你一个字符串 s ，请找出满足每个字符最多出现两次的最长子字符串，并返回该 子字符串 的 最大 长度。

func main() {
	fmt.Println(maximumLengthSubstring("bcbbbcba"))
}

// 输入： s = "bcbbbcba"
// 输出： 4
// 解释：
// 以下子字符串长度为 4，并且每个字符最多出现两次："bcba"。

//输入： s = "aaaa"
//输出： 2
//解释：
//以下子字符串长度为 2，并且每个字符最多出现两次："aaaa"。

func maximumLengthSubstring(s string) int {
	mp := [25]int{}
	left := 0
	ans := 0
	for i := 0; i < len(s); i++ {
		mp[s[i]-'a']++
		// 每个字符最多出现两次
		for mp[s[i]-'a'] > 2 {
			// 左边界
			mp[s[left]-'a']--
			left++
		}
		if i-left+1 > ans {
			ans = i - left + 1
		}
	}
	return ans
}
