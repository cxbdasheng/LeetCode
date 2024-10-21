package main

import "fmt"

// 3 3
// 1 2 1
// 1 1 1
// 1 1 3
func main() {
	var m, n int
	fmt.Scan(&n, &m)
	if n < 2 || m < 2 {
		fmt.Println(0)
		return
	}
	nums := make([][]int, n)
	for i := 0; i < n; i++ {
		nums[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&nums[i][j])
		}
	}
	top := 0
	left := 0
	right := len(nums[0]) - 1
	bottom := len(nums) - 1
	var res int
	for left < right || top < bottom {
		left = 0
		for ; left <= right-1; left++ {
			if nums[top][left] == nums[top][left+1] && nums[top][left+1] == nums[top+1][left] && nums[top+1][left] == nums[top+1][left+1] {
				res++
			}
		}
		top++
	}
	fmt.Println(res)
}
