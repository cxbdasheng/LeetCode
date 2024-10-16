package main

import "fmt"

// 12
// 3 9 3 2 5 6 7 3 2 3 3 3
func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	mp := make(map[int]int)
	mp[0] = 0
	max := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
		mp[nums[i]]++
		if mp[nums[i]] >= n/2 {
			if mp[max] < mp[nums[i]] {
				max = nums[i]
			}
		}
	}
	fmt.Println(max)
	max(1, 2)
}
