package main

import (
	"fmt"
	"strings"
)

// 给你一个长度为 n 下标从 0 开始的字符串 blocks ，blocks[i] 要么是 'W' 要么是 'B' ，表示第 i 块的颜色。字符 'W' 和 'B' 分别表示白色和黑色。
//
// 给你一个整数 k ，表示想要 连续 黑色块的数目。
//
// 每一次操作中，你可以选择一个白色块将它 涂成 黑色块。
//
// 请你返回至少出现 一次 连续 k 个黑色块的 最少 操作次数。
func main() {
	fmt.Println(minimumRecolors("WBBWWBBWBW", 7))
}
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
