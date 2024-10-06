package main

import "fmt"

func main() {
	fmt.Println(wiggleMaxLength([]int{1, 7, 4, 9, 2, 5}))
}
func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	if len(nums) == 2 && nums[0] == nums[1] {
		return 1
	}
	//var result []int
	//sum := 1
	//for i := 1; i < len(nums); i++ {
	//	if i == 1 {
	//		if nums[i] == nums[i-1] {
	//			continue
	//		}
	//		sum++
	//		result = append(result, nums[i]-nums[i-1])
	//		continue
	//	}
	//	if (nums[i]-nums[i-1] > 0 && result[len(result)-1] < 0) || (nums[i]-nums[i-1] < 0 && result[len(result)-1] > 0) {
	//		sum++
	//		result = append(result, nums[i]-nums[i-1])
	//	}
	//
	//}
	//return sum
	prediff := 0
	curdiff := 0
	res := 1
	for i := 0; i < len(nums)-1; i++ {
		curdiff = nums[i+1] - nums[i]
		if (prediff >= 0 && curdiff < 0) || (prediff <= 0 && curdiff > 0) {
			res++
			prediff = curdiff
		}
	}
	return res
}
