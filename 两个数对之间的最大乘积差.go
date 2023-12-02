package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{5, 6, 2, 7, 4}
	sort.Ints(nums)
	fmt.Println((nums[len(nums)-1] * nums[len(nums)-2]) - (nums[0] * nums[1]))
}
