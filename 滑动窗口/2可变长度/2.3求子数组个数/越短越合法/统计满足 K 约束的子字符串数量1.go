package main

import "fmt"

// 给你一个 二进制 字符串 s 和一个整数 k。
// 如果一个 二进制字符串 满足以下任一条件，则认为该字符串满足 k 约束：
// 字符串中 0 的数量最多为 k。
// 字符串中 1 的数量最多为 k。
// 返回一个整数，表示 s 的所有满足 k 约束 的 子字符串 的数量。
func main() {
	fmt.Println(countKConstraintSubstrings("10101", 1))
}

// 输入：s = "10101", k = 1
// 输出：12
// 解释：
// s 的所有子字符串中，除了 "1010"、"10101" 和 "0101" 外，其余子字符串都满足 k 约束。

// 输入：s = "1010101", k = 2
// 输出：25
// 解释：
// s 的所有子字符串中，除了长度大于 5 的子字符串外，其余子字符串都满足 k 约束。

// 输入：s = "11111", k = 1
// 输出：15
// 解释：
// s 的所有子字符串都满足 k 约束。

func countKConstraintSubstrings(s string, k int) int {
	var ans int
	l := 0
	n := len(s)
	mp := make([]int, 2)
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			mp[1]++
		} else {
			mp[0]++
		}
		for l <= i && mp[0] > k && mp[1] > k {
			if s[l] == '1' {
				mp[1]--
			} else {
				mp[0]--
			}
			l++
		}
		ans += i - l + 1
	}
	return ans
}
