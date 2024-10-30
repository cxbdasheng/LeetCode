package main

import (
	"fmt"
	"math"
)

// 给你两个长度相同的字符串，s 和 t。
// 将 s 中的第 i 个字符变到 t 中的第 i 个字符需要 |s[i] - t[i]| 的开销（开销可能为 0），也就是两个字符的 ASCII 码值的差的绝对值。
// 用于变更字符串的最大预算是 maxCost。在转化字符串时，总开销应当小于等于该预算，这也意味着字符串的转化可能是不完全的。
// 如果你可以将 s 的子字符串转化为它在 t 中对应的子字符串，则返回可以转化的最大长度。
// 如果 s 中没有子字符串可以转化成 t 中对应的子字符串，则返回 0。

func main() {
	fmt.Println(equalSubstring("abcd", "bcdf", 3))
}

// 输入：s = "abcd", t = "bcdf", maxCost = 3
// 输出：3
// 解释：s 中的 "abc" 可以变为 "bcd"。开销为 3，所以最大长度为 3。

// 输入：s = "abcd", t = "cdef", maxCost = 3
// 输出：1
// 解释：s 中的任一字符要想变成 t 中对应的字符，其开销都是 2。因此，最大长度为 1。

// 输入：s = "abcd", t = "acde", maxCost = 0
// 输出：1
// 解释：a -> a, cost = 0，字符串未发生变化，所以最大长度为 1。

// 1 <= s.length, t.length <= 10^5
// 0 <= maxCost <= 10^6
// s 和 t 都只含小写英文字母。
func equalSubstring(s string, t string, maxCost int) int {
	left := 0
	m := 0
	sum := maxCost
	for i := 0; i < len(s); i++ {
		sum -= int(math.Abs(float64(s[i]) - float64(t[i])))
		for sum < 0 {
			if left > i {
				sum = maxCost
				break
			}
			sum += int(math.Abs(float64(s[left]) - float64(t[left])))
			left++
		}

		if m < i-left+1 {
			m = i - left + 1
		}
	}
	return m
}
