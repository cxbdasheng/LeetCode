package main

import "fmt"

// 给你一个二进制字符串 s 和一个整数 k 。如果所有长度为 k 的二进制字符串都是 s 的子串，请返回 true ，否则请返回 false 。
func main() {
	fmt.Println(hasAllCodes("00110110", 2))
}

// 输入：s = "00110110", k = 2
// 输出：true
// 解释：长度为 2 的二进制串包括 "00"，"01"，"10" 和 "11"。它们分别是 s 中下标为 0，1，3，2 开始的长度为 2 的子串。

// 输入：s = "0110", k = 1
// 输出：true
// 解释：长度为 1 的二进制串包括 "0" 和 "1"，显然它们都是 s 的子串。

// 输入：s = "0110", k = 2
// 输出：false
// 解释：长度为 2 的二进制串 "00" 没有出现在 s 中。
func hasAllCodes(s string, k int) bool {
	if len(s) < k {
		return false
	}
	mp := make(map[string]bool)
	mp[s[:k]] = true
	for i := k + 1; i <= len(s); i++ {
		mp[s[i-k:i]] = true
	}
	// 1<<k 长度为 k 的二进制串的数量为 2k
	// 所有长度为 k 的二进制字符串都是 s 的子串，也就是说 len(mp) == 1<<k
	return len(mp) == 1<<k
}
