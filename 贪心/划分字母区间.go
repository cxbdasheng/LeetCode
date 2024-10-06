package main

import "fmt"

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}
func partitionLabels(s string) []int {
	var res []int
	mp := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		mp[s[i]] = i
	}
	left := 0
	rigth := 0
	for i := 0; i < len(s); i++ {
		rigth = Max(rigth, mp[s[i]])
		if rigth == i {
			res = append(res, rigth-left+1)
			left = rigth + 1
		}
	}

	return res
}
func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
