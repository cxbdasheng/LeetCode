package main

import "fmt"

func main() {
	fmt.Println(strStr("aabaabaaf", "aabaaf"))
}
func strStr(haystack string, needle string) int {
	next := getNext(make([]int, len(needle)), needle)
	j := 0
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		// 这里是后加
		if needle[j] == haystack[i] {
			j++
		}
		if j == len(needle) {
			return i - j + 1
		}
	}
	return -1
}

func getNext(next []int, s string) []int {
	// j 为前缀开头，i为后缀末尾。
	j := 0
	for i := 1; i < len(s); i++ {
		// j > 0 && s[j] != s[i] 为灵魂语句；
		// j > 0 1.防止死循环；2.防止越界报错；
		// for s[j] != s[i] 可以不断为匹配回溯过程
		for j > 0 && s[j] != s[i] {
			j = next[j-1]
		}
		if s[j] == s[i] {
			j++
		}
		next[i] = j
	}
	return next
}
