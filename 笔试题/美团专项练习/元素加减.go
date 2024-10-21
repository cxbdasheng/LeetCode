package main

import (
	"fmt"
	"sort"
)

// 4 5
// -2 3 -2 9
func main() {
	var k, n int
	fmt.Scan(&n, &k)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
		if nums[i] < 0 {
			nums[i] *= -1
		}
	}
	sort.Ints(nums)
	var res int
	for i := 0; i < n; i++ {
		if nums[i] < k {
			k -= nums[i]
			res++
		}
	}
	fmt.Println(res)
}
