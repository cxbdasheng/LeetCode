package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
	str := "azsxfq"
	s := []rune(str)
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	fmt.Println(string(s))
}
func majorityElement(nums []int) int {
	max := 0
	n := 0
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		if m[nums[i]] >= len(nums)/2 && m[nums[i]] > max {
			max = m[nums[i]]
			n = nums[i]
		}
	}
	return n
}
