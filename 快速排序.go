package main

import "fmt"

func main() {
	var nums = []int{5, 10, 8, 3, 555, 2, 6, 4, 7, 1, 9}

	fmt.Println(quickSort(nums))
}

func quickSort(nums []int) []int {
	//if l < r {
	//	pivot := r
	//	counter := l
	//	for i := l; i < r; i++ {
	//		if nums[i] < nums[pivot] {
	//			nums[i], nums[counter] = nums[counter], nums[i]
	//			counter++
	//		}
	//	}
	//	nums[pivot], nums[counter] = nums[counter], nums[pivot]
	//	quickSort(nums, l, counter-1)
	//	quickSort(nums, counter+1, r)
	//}
	if len(nums) < 2 {
		return nums
	}
	pivot := nums[0]
	var left, rigth []int
	for _, ele := range nums[1:] {
		if ele < pivot {
			left = append(left, ele)
		} else {
			rigth = append(rigth, ele)
		}
	}
	return append(quickSort(left), append([]int{pivot}, quickSort(rigth)...)...)
}
