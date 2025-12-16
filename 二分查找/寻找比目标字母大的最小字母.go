package main

import (
	"fmt"
)

func main() {
	fmt.Println(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'c'))
}

func nextGreatestLetter(letters []byte, target byte) byte {
	left := 0
	right := len(letters) - 1
	if target >= letters[right] {
		return letters[0]
	}
	target = target + 1
	// 标准二分查找
	for left <= right {

		mid := (left + right) >> 1
		if letters[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return letters[left]

}
