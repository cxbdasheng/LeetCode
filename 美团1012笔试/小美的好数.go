package main

import "fmt"

// 3
// 13 14 15
func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	var res int
	for i := 0; i < n; i++ {
		if IsNice(nums[i]) {
			res++
		}
	}
	fmt.Println(res)
}

func IsNice(num int) bool {
	var sum, c int
	old := false
	for num != 0 {
		c = num % 10
		if c&1 == 1 {
			old = true
		}
		sum += c
		num /= 10
	}
	if sum&1 != 1 && old {
		return true
	}
	return false
}
