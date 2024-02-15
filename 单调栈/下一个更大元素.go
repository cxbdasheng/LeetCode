package main

import "fmt"

func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}
	var stack []int
	result := make([]int, len(nums1))
	for i := range result {
		result[i] = -1
	}
	nums1Map := make(map[int]int, len(nums1))
	for k, v := range nums1 {
		nums1Map[v] = k
	}
	stack = append(stack, 0)
	for i := 1; i < len(nums2); i++ {
		top := stack[len(stack)-1]
		if nums2[top] >= nums2[i] {
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && nums2[top] < nums2[i] {
				if v, ok := nums1Map[nums2[top]]; ok {
					result[v] = nums2[i]
				}
				stack = stack[:len(stack)-1]
				if len(stack) > 0 {
					top = stack[len(stack)-1]
				}
			}
			stack = append(stack, i)
		}
	}

	return result
}
