package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{2, 0}))
}

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	cover := 0
	for i := 0; i <= cover; i++ {
		cover = Max(cover, nums[i]+i)
		if cover >= len(nums)-1 {
			return true
		}
	}
	return false

}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
