package main

import "fmt"

// 给你一个字符串 s ，请你返回满足以下条件且出现次数最大的 任意 子串的出现次数：
//
//	子串中不同字母的数目必须小于等于 maxLetters 。
//	子串的长度必须大于等于 minSize 且小于等于 maxSize 。
//
// 输入：s = "aababcaab", maxLetters = 2, minSize = 3, maxSize = 4
// 输出：2
// 解释：子串 "aab" 在原字符串中出现了 2 次。
// 它满足所有的要求：2 个不同的字母，长度为 3 （在 minSize 和 maxSize 范围内）。

// 本题的重中之重（也是我卡的最久的地方）：本题只需统计长度为minSize的子串，而不需要统计长度为maxSize的子串。
// Why?
// "abc" 肯定会覆盖 a，ab， 即长的肯定会覆盖短的，只要考虑最短的就好咯。
func main() {
	fmt.Println(maxFreq("aaaaacbc", 2, 4, 6))
}

// 输入：s = "aababcaab", maxLetters = 2, minSize = 3, maxSize = 4
// 输出：2
// 解释：子串 "aab" 在原字符串中出现了 2 次。
// 它满足所有的要求：2 个不同的字母，长度为 3 （在 minSize 和 maxSize 范围内）。

// 输入：s = "aaaa", maxLetters = 1, minSize = 3, maxSize = 3
// 输出：2
// 解释：子串 "aaa" 在原字符串中出现了 2 次，且它们有重叠部分。

// 输入：s = "aabcabcab", maxLetters = 2, minSize = 2, maxSize = 3
// 输出：3
func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	mp := make(map[string]int)
	if validate(s[:minSize], maxLetters) {
		mp[s[:minSize]] = 1
	}
	// 字符串一定要注意边界值取值
	for i := minSize + 1; i <= len(s); i++ {
		if validate(s[i-minSize:i], maxLetters) {
			mp[s[i-minSize:i]]++
		}
	}
	m := 0
	for _, v := range mp {
		if v > m {
			m = v
		}
	}
	return m

}

func validate(s string, l int) bool {
	mp := make(map[string]int)
	for i := 0; i < len(s); i++ {
		mp[string(s[i])]++
	}
	return len(mp) <= l
}
