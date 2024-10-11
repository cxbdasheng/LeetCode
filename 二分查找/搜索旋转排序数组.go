package main

import "fmt"

func main() {
	fmt.Println(search([]int{3, 1}, 3))
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > nums[left] {
			if nums[left] <= target && target < nums[mid] {
				right = mid
			} else {
				left = mid + 1
			}
		} else {
			if nums[right-1] >= target && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}
	return -1
}
