package main

import "fmt"

// 3
// 13 14 15
func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	var res int
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {

		}
	}
	fmt.Println(res)
}
