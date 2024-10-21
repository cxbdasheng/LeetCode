package main

import (
	"fmt"
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
				nums = strToInt(s[j:i], nums)
				j = i
			}
			j++
		}
		if len(s)-1 == i && i != j {
			nums = strToInt(s[j:i], nums)
		}
	}

	if k > len(nums) {
		println('N')
	} else {
		println(nums[k-1])
	}
}

func strToInt(str string, nums []int) []int {
	if len(str) == 1 && str[0] == '0' {
		nums = append(nums, 0)
		return nums
	}
	base := 1
	res := 0
	for i := len(str) - 1; i >= 0; i-- {
		res += int(str[i]-'0') * base
		base *= 10
	}
	nums = append(nums, res)
	return nums
}
