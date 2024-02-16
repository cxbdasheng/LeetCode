package main

import "fmt"

func main() {
	fmt.Println(nextGreaterElements([]int{1, 2, 1}))
}

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}
	stack := []int{0}
	result := make([]int, n)
	for i, _ := range result {
		result[i] = -1
	}
	for i := 1; i < n*2; i++ {
		j := i % n
		top := stack[len(stack)-1]
		if nums[j] <= nums[top] {
			stack = append(stack, j)
		} else {
			// len(stack) > 0 防止死循环
			for len(stack) > 0 && nums[j] > nums[top] {
				result[top] = nums[j]
				stack = stack[:len(stack)-1]
				if len(stack) > 0 {
					top = stack[len(stack)-1]
				}
			}
			stack = append(stack, j)

		}
	}
	return result
}
