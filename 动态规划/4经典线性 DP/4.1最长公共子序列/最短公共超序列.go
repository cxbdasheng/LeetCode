package main

import "fmt"

func main() {
	fmt.Println(shortestCommonSupersequence("abac", "cab"))
}
func shortestCommonSupersequence(str1 string, str2 string) string {
	if str1 == "" { // str1 是空串，返回剩余的 str2
		return str2
	}
	if str2 == "" { // str2 是空串，返回剩余的 str1
		return str1
	}
	n, m := len(str1), len(str2)
	if str1[n-1] == str2[m-1] { // 最短公共超序列一定包含 str1[n-1]
		return shortestCommonSupersequence(str1[:n-1], str2[:m-1]) + str1[n-1:]
	}

	ans1 := shortestCommonSupersequence(str1[:n-1], str2)
	ans2 := shortestCommonSupersequence(str1, str2[:m-1])

	if len(ans1) < len(ans2) { // 取 ans1 和 ans2 中更短的组成答案
		return ans1 + str1[n-1:]
	}
	return ans2 + str2[m-1:]
}
