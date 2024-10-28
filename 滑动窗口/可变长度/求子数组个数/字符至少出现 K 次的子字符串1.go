package main

import "fmt"

func main() {
	fmt.Println(numberOfSubstrings("abacb", 2))
}
func numberOfSubstrings(s string, k int) int {
	ans := 0
	left := 0
	mp := make([]int, 26)
	// 当前进窗口的数会大于 k 则继续执行
	for i := 0; i < len(s); i++ {
		mp[s[i]-'a']++
		for mp[s[i]-'a'] >= k {
			ans += len(s) - i
			mp[s[left]-'a']--
			left++
		}
	}
	return ans
}
