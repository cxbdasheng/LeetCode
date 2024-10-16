package main

import "fmt"

// 4
// 1 3 2 4
// WRRW
func main() {
	var n int
	fmt.Scan(&n)
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	var str string
	fmt.Scan(&str)

	color := make([]int, n)
	for i := 0; i < n; i++ {
		if str[i] == 'W' {
			color[i] = 0
		} else {
			color[i] = 1
		}
	}
	var result int
	mp := make(map[int]int)
	for i := 0; i < n; i++ {
		mp[arr[i]] = color[i]
	}
	for i := 0; i < n; i++ {
		if arr[i] == i+1 {
			continue
		}
		if arr[i] != i+1 && color[i] != 0 {

		}
	}
}
