package main

import "fmt"

// 6
// 1 1 4 5 1 4
func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	var p, o int
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
		if i&1 == 1 && o < nums[i] {
			o = nums[i]
		}
		if i&1 != 1 && p < nums[i] {
			p = nums[i]
		}
	}
	sum := 0
	for i := 0; i < n; i++ {
		if i&1 == 1 {
			sum += o - nums[i]
		} else {
			sum += p - nums[i]
		}
	}
	if p != o {
		fmt.Println(sum)
	} else {
		fmt.Println(sum + n/2)
	}
}
