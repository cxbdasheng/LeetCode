package main

import "fmt"

func main() {
	fmt.Println(getSubarrayBeauty([]int{1, -1, -3, -2, 3}, 3, 2))
}
func getSubarrayBeauty(nums []int, k int, x int) []int {
	var stack []int
	for i := 0; i < k; i++ {
		if len(stack) == 0 || nums[i] < 0 {
			stack = append(stack, i)
		} else {
			top := len(stack) - 1
			for len(stack) > 0 && nums[top] > nums[i] {
				stack = stack[:top]
				if len(stack) > 0 {
					top = stack[len(stack)-1]
				}
			}
		}
	}
	return []int{1}
}
