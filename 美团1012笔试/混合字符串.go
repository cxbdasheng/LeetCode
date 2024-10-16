package main

import (
	"fmt"
	"strconv"
)

// 3
// a12b03c3d0
func main() {

	var k int
	fmt.Scan(&k)
	var s string
	fmt.Scan(&s)
	var nums []int
	j := 0
	i := 0
	for ; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			if i != j {
				temp, _ := strconv.Atoi(s[j:i])
				nums = append(nums, temp)
				j = i + 1
			} else {
				j++
			}
		}
	}

	if i != j {
		temp, _ := strconv.Atoi(s[j:i])
		nums = append(nums, temp)
	}
	if len(nums) > k-1 {
		fmt.Println(nums[k-1])
	} else {
		fmt.Println(string('N'))
	}
}
