package main

import "fmt"

// 给你一个由字符 'a'、'b'、'c' 组成的字符串 s 和一个非负整数 k 。每分钟，你可以选择取走 s 最左侧 还是 最右侧 的那个字符。
// 你必须取走每种字符 至少 k 个，返回需要的 最少 分钟数；如果无法取到，则返回 -1 。
func main() {
	fmt.Println(takeCharacters("abc", 1))
}

// 输入：s = "aabaaaacaabc", k = 2
// 输出：8
// 解释：
// 从 s 的左侧取三个字符，现在共取到两个字符 'a' 、一个字符 'b' 。
// 从 s 的右侧取五个字符，现在共取到四个字符 'a' 、两个字符 'b' 和两个字符 'c' 。
// 共需要 3 + 5 = 8 分钟。
// 可以证明需要的最少分钟数是 8 。

//输入：s = "a", k = 1
//输出：-1
//解释：无法取到一个字符 'b' 或者 'c'，所以返回 -1 。

func takeCharacters(s string, k int) int {
	cnt := [3]int{}
	for _, c := range s {
		cnt[c-'a']++ // 一开始，把所有字母都取走
	}
	if cnt[0] < k || cnt[1] < k || cnt[2] < k {
		return -1 // 字母个数不足 k
	}
	mx, left := 0, 0
	for right, c := range s {
		c -= 'a'
		cnt[c]--         // 移入窗口，相当于不取走 c
		for cnt[c] < k { // 窗口之外的 c 不足 k
			cnt[s[left]-'a']++ // 移出窗口，相当于取走 s[left]
			left++
		}
		mx = max(mx, right-left+1)
	}
	return len(s) - mx
}
