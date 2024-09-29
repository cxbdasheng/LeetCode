package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	area := 0
	for left < right {
		mn := Min(height[left], height[right])
		area = Max(area, (right-left)*mn)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return area
}

func Min(left int, right int) int {
	if left > right {
		return right
	}
	return left
}

func Max(left int, right int) int {
	if left > right {
		return left
	}
	return right
}
