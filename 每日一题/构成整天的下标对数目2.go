package main

import "fmt"

func main() {
	fmt.Println(countCompleteDayPairs([]int{12, 12, 30, 24, 24}))
}
func countCompleteDayPairs(hours []int) int64 {
	mp := make(map[int]int64, 24)
	var ans int64
	for _, h := range hours {
		ans += mp[(24-h%24)%24]
		mp[h%24]++
	}
	return ans
}
