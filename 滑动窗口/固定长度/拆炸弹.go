package main

import "fmt"

// 你有一个炸弹需要拆除，时间紧迫！你的情报员会给你一个长度为 n 的 循环 数组 code 以及一个密钥 k 。
//
// 为了获得正确的密码，你需要替换掉每一个数字。所有数字会 同时 被替换。
//
//	如果 k > 0 ，将第 i 个数字用 接下来 k 个数字之和替换。
//	如果 k < 0 ，将第 i 个数字用 之前 k 个数字之和替换。
//	如果 k == 0 ，将第 i 个数字用 0 替换。
//
// 由于 code 是循环的， code[n-1] 下一个元素是 code[0] ，且 code[0] 前一个元素是 code[n-1] 。
//
// 给你 循环 数组 code 和整数密钥 k ，请你返回解密后的结果来拆除炸弹！
func main() {
	fmt.Println(decrypt([]int{5, 7, 1, 4}, 3))
}
func decrypt(code []int, k int) []int {
	n := len(code)
	ans := make([]int, len(code))
	sum := 0
	//因为还要隔一个
	r := k + 1
	if k < 0 {
		r = n
		k = -k
	}
	for _, v := range code[r-k : r] {
		sum += v
	}
	for i := 0; i < n; i++ {
		ans[i] = sum
		sum += code[r%n] - code[(r-k)%n]
		r++
	}
	return ans
}
