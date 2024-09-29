package main

import "fmt"

func main() {
	fmt.Println(sub([]int{1, 3, 7, 5, 2}))
}

func sub(nums []int) []int {
	d := make([]int, len(nums)+1)
	add(d, 2, 4, 5)
	add(d, 1, 3, 2)
	add(d, 0, 2, -3)

	for i := 1; i < len(nums); i++ {
		d[i] += d[i-1]
	}
	fmt.Println(d)
	for i := 0; i < len(nums); i++ {
		nums[i] += d[i]
	}

	return nums
}

func add(d []int, l int, r int, v int) {
	d[l] = d[l] + v
	d[r+1] = d[r+1] - v
}
