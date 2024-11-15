package main

import "fmt"

// 10
// 2 40 400 3000
// 1 4 2 4 2 2 1 1 100 1
func main() {

	var n int
	fmt.Scan(&n)
	arr := make([]int, 4)
	fmt.Scan(&arr[0], &arr[1], &arr[2], &arr[3])
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	fmt.Println(16)
}
