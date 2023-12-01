package main

import "fmt"

func main() {
	var nums = []int{4, 1, 2, 1, 2}
	var result int
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			result = nums[i]
			continue
		}
		result = result ^ nums[i]
	}
	fmt.Println(result)
}
