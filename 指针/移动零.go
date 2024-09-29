package main

import "fmt"

func main() {
	fmt.Println(moveZeroes([]int{2, 1, 0, 3, 12}))
}
func moveZeroes(nums []int) []int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return nums
}
