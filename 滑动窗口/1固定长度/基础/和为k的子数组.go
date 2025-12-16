package main

import (
	"fmt"
)

func main() {
	fmt.Println(subarraySum([]int{1, 2, 1, 2, 1}, 3))
}
func subarraySum(nums []int, k int) int {
	cnt := make(map[int]int, len(nums)+1)
	cnt[0] = 1
	s := 0
	ans := 0
	for _, num := range nums {
		s += num
		ans += cnt[s-k]
		cnt[s]++
	}
	return ans
}
