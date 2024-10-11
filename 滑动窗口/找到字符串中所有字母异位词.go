package main

import (
	"fmt"
)

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
}
func findAnagrams(s string, p string) []int {
	var res []int
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return res
	}

	var sCount, pCount [26]int
	for i, ch := range p {
		sCount[s[i]-'a']++
		pCount[ch-'a']++
	}
	if sCount == pCount {
		res = append(res, 0)
	}

	for i, ch := range s[:sLen-pLen] {
		sCount[ch-'a']--
		sCount[s[i+pLen]-'a']++
		if sCount == pCount {
			res = append(res, i+1)
		}
	}

	return res
}
