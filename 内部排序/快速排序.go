package main

import "fmt"

func main() {
	fmt.Println(fastSort([]int{68, 13, 58, 95, 10, 57, 62, 13, 106, 78, 23, 85}))
}

func fastSort(nums []int) []int {
	if len(nums) < 2 {
		return nums // 基本情况：空数组或单元素数组已排序
	}

	pivot := nums[0]
	var left []int
	var right []int
	var equal []int

	for _, value := range nums {
		switch {
		case value < pivot:
			left = append(left, value)
			break
		case value > pivot:
			right = append(right, value)
			break
		default:
			equal = append(equal, value)
			break
		}
	}

	return append(append(fastSort(left), equal...), fastSort(right)...)
}
