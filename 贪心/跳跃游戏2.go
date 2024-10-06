package main

import (
	"fmt"
)

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}

func jump(nums []int) int {
	if len(nums) < 3 {
		return len(nums) - 1
	}
	res := 0
	cover := 0
	next := 0

	for i := 0; i < len(nums); i++ {
		next = Max(i+nums[i], next)
		if i == cover && i != len(nums)-1 {
			cover = next
			res++
		}
	}

	return res
	//dp := make([]int, len(nums))
	//dp[0] = 0
	//dp[1] = 1
	//for i := 2; i < len(nums); i++ {
	//	dp[i] = math.MaxInt
	//	for j := i - 1; j >= 0; j-- {
	//		if j+nums[j] >= i {
	//			dp[i] = Min(dp[i], dp[j]+1)
	//		}
	//	}
	//}
	//return dp[len(nums)-1]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
