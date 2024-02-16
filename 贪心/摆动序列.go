package main

import "fmt"

func main() {
	fmt.Println(wiggleMaxLength([]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}))
}
func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	var result []int
	sum := 1
	for i := 1; i < len(nums); i++ {
		if i == 1 {
			if nums[i] == nums[i-1] {
				continue
			}
			sum++
			result = append(result, nums[i]-nums[i-1])
			continue
		}
		if (nums[i]-nums[i-1] > 0 && result[len(result)-1] < 0) || (nums[i]-nums[i-1] < 0 && result[len(result)-1] > 0) {
			sum++
			result = append(result, nums[i]-nums[i-1])
		}

	}
	return sum
}
