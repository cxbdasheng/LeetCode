package main

import "fmt"

// 有一个只含有 'Q', 'W', 'E', 'R' 四种字符，且长度为 n 的字符串。
// 假如在该字符串中，这四个字符都恰好出现 n/4 次，那么它就是一个「平衡字符串」。
// 给你一个这样的字符串 s，请通过「替换一个子串」的方式，使原字符串 s 变成一个「平衡字符串」。
// 你可以用和「待替换子串」长度相同的 任何 其他字符串来完成替换。
// 请返回待替换子串的最小可能长度。
//
// 如果原字符串自身就是一个平衡字符串，则返回 0。
func main() {
	fmt.Println(balancedString("WWEQERQWQWWRWWERQWEQ"))
}

//输入：s = "QWER"
//输出：0
//解释：s 已经是平衡的了。

//输入：s = "QQWE"
//输出：1
//解释：我们需要把一个 'Q' 替换成 'R'，这样得到的 "RQWE" (或 "QRWE") 是平衡的。

// 输入：s = "QQQW"
// 输出：2
// 解释：我们可以把前面的 "QQ" 替换成 "ER"。

// 输入：s = "QQQQ"
// 输出：3
// 解释：我们可以替换后 3 个 'Q'，使 s = "QWER"。
func balancedString(s string) int {
	cnt, m := [4]int{}, len(s)/4 // 也可以用哈希表，不过数组更快一些
	for i := 0; i < len(s); i++ {
		cnt[GetIndex(s[i])]++
	}
	if cnt[0] == m && cnt[1] == m && cnt[2] == m && cnt[3] == m {
		return 0 // 已经符合要求啦
	}
	ans, left := len(s), 0
	for i := 0; i < len(s); i++ {
		cnt[GetIndex(s[i])]--
		// 判断外面的是否符合要求
		for cnt[0] <= m && cnt[1] <= m && cnt[2] <= m && cnt[3] <= m {
			ans = min(ans, i-left+1)
			cnt[GetIndex(s[left])]++
			left++ // 缩小子串
		}
	}
	return ans
}
func GetIndex(s byte) int {
	ans := 0
	if s == 'W' {
		ans = 1
	} else if s == 'E' {
		ans = 2
	} else if s == 'R' {
		ans = 3
	}
	return ans
}
