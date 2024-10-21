package main

import "fmt"

// 3
// 5
// 3 6 2 3 1
// 6
// 1 1 4 5 1 4
// 10
// 7 8 3 5 6 1 2 4 9 10
func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var m int
		var lenght int = 1
		fmt.Scan(&m)
		nums := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&nums[j])
		}
		dp := make([]int, m)
		dp[0] = 1
		a := 0
		cur := 0
		for j := 1; j < m; j++ {
			if nums[j] > nums[j-1] {
				lenght++
				if lenght > a {
					a = lenght
					cur = j
				}
				dp[j] = max(dp[j-1], lenght)
			} else {
				lenght = 1
				dp[j] = dp[j-1]
			}
		}
		ma := dp[m-1]
		if 0 < cur && cur < m && nums[cur-ma] > nums[cur-ma-1] {
			ma = ma + dp[cur-ma-1] + 1
		} else {
			ma = ma + 1
		}
		fmt.Println(ma)
	}
}
