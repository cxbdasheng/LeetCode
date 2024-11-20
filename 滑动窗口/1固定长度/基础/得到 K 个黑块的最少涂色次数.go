package main

import (
	"fmt"
	"strings"
)

// 给你一个长度为 n 下标从 0 开始的字符串 blocks ，blocks[i] 要么是 'W' 要么是 'B' ，表示第 i 块的颜色。字符 'W' 和 'B' 分别表示白色和黑色。
// 给你一个整数 k ，表示想要 连续 黑色块的数目。
// 每一次操作中，你可以选择一个白色块将它 涂成 黑色块。
//
// 请你返回至少出现 一次 连续 k 个黑色块的 最少 操作次数。
func main() {
	fmt.Println(minimumRecolors("WBBWWBBWBW", 7))
}

// 输入：blocks = "WBBWWBBWBW", k = 7
// 输出：3
// 解释：
// 一种得到 7 个连续黑色块的方法是把第 0 ，3 和 4 个块涂成黑色。
// 得到 blocks = "BBBBBBBWBW" 。
// 可以证明无法用少于 3 次操作得到 7 个连续的黑块。
// 所以我们返回 3 。

// 输入：blocks = "WBWBBBW", k = 2
// 输出：0
// 解释：
// 不需要任何操作，因为已经有 2 个连续的黑块。
// 所以我们返回 0 。
func minimumRecolors(blocks string, k int) int {
	//sum := 0
	//m := 0
	//for i := 0; i < k; i++ {
	//	if blocks[i] == 'B' {
	//		sum += 1
	//	}
	//}
	//if sum == k {
	//	return 0
	//}
	//m = sum
	//for i := k; i < len(blocks); i++ {
	//	if blocks[i] == 'B' {
	//		sum++
	//	}
	//	if blocks[i-k] == 'B' {
	//		sum--
	//	}
	//	if m < sum {
	//		m = sum
	//	}
	//}
	//if m == k {
	//	return 0
	//} else if m < k {
	//	return k - m
	//}
	//return 0

	cntW := strings.Count(blocks[:k], "W")
	ans := cntW
	for i := k; i < len(blocks); i++ {
		cntW += int(blocks[i]&1) - int(blocks[i-k]&1)
		ans = min(ans, cntW)
	}
	return ans
}
