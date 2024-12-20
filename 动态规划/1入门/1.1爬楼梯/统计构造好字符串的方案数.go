package main

import "fmt"

const M = 1_000_000_007

// 给你整数 zero ，one ，low 和 high ，我们从空字符串开始构造一个字符串，每一步执行下面操作中的一种：
// 将 '0' 在字符串末尾添加 zero  次。
// 将 '1' 在字符串末尾添加 one 次。

// 以上操作可以执行任意次。
// 如果通过以上过程得到一个 长度 在 low 和 high 之间（包含上下边界）的字符串，那么这个字符串我们称为 好 字符串。
// 请你返回满足以上要求的 不同 好字符串数目。由于答案可能很大，请将结果对 109 + 7 取余 后返回。

func main() {
	fmt.Println(countGoodStrings(3, 3, 2, 1))
}

// 输入：low = 3, high = 3, zero = 1, one = 1
// 输出：8
// 解释：
// 一个可能的好字符串是 "011" 。
// 可以这样构造得到："" -> "0" -> "01" -> "011" 。
// 从 "000" 到 "111" 之间所有的二进制字符串都是好字符串。

// 输入：low = 2, high = 3, zero = 1, one = 2
// 输出：5
// 解释：好字符串为 "00" ，"11" ，"000" ，"110" 和 "011" 。
func countGoodStrings(low int, high int, zero int, one int) int {
	var ans int
	dp := make([]int, high+1)
	dp[0] = 1

	for i := 1; i <= high; i++ {
		// 如果 i≥zero，那么可以在长为 i−zero 的字符串末尾添加 zero 个 0，方案数为 f[i−zero]
		if i >= zero {
			dp[i] = (dp[i] + dp[i-zero]) % M
		}
		// 如果 i≥one，那么可以在长为 i−one 的字符串末尾添加 one 个 1，方案数为 f[i−one]
		if i >= one {
			// 两类方案互斥（第 i 个字符不能既是 0 又是 1），所以用加法原理，得
			dp[i] = (dp[i] + dp[i-one]) % M
		}
		if i >= low {
			ans = (dp[i] + ans) % M
		}
	}
	return ans
}
