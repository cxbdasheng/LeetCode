package main

import (
	"fmt"
	"slices"
)

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}
func rotate(nums []int, k int) {
	k = k % len(nums)
	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
	fmt.Println(nums)
}
