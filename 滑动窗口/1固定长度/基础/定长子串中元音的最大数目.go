package main

import "fmt"

// 给你字符串 s 和整数 k 。
// 请返回字符串 s 中长度为 k 的单个子字符串中可能包含的最大元音字母数。
// 英文中的 元音字母 为（a, e, i, o, u）。
func main() {
	fmt.Println(maxVowels("abciiidef", 3))
}

//输入：s = "abciiidef", k = 3
//输出：3
//解释：子字符串 "iii" 包含 3 个元音字母。

//输入：s = "aeiou", k = 2
//输出：2
//解释：任意长度为 2 的子字符串都包含 2 个元音字母。

// 输入：s = "leetcode", k = 3
// 输出：2
// 解释："lee"、"eet" 和 "ode" 都包含 2 个元音字母。

//输入：s = "rhythms", k = 4
//输出：0
//解释：字符串 s 中不含任何元音字母。

func maxVowels(s string, k int) int {
	m := 0
	vowel := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			vowel++
			if vowel > m {
				m = vowel
			}
		}
		if i+1 < k {
			continue
		}
		if s[i+1-k] == 'a' || s[i+1-k] == 'e' || s[i+1-k] == 'i' || s[i+1-k] == 'o' || s[i+1-k] == 'u' {
			vowel--
		}
	}
	return m
}
