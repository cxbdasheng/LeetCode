package main

import "fmt"

const M = int64(1e9 + 7)

// 4
// 1 2 3 4
func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	flag := 1
	res := backtrack(nums, flag)
	fmt.Println(res[0] % M)
}

func backtrack(nums []int64, flag int) []int64 {
	if len(nums) == 1 {
		return nums
	}
	var res []int64
	for i := 1; i < len(nums); i++ {
		if flag == 1 {
			res = append(res, nums[i]+nums[i-1])
		} else {
			res = append(res, nums[i-1]-nums[i])
		}
		flag *= -1
	}
	return backtrack(res, flag)
}
