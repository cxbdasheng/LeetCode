package main

import "fmt"

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}
func longestConsecutive(nums []int) int {
	hashMap := make(map[int]bool)
	for _, num := range nums {
		hashMap[num] = true
	}
	max := 0
	for num, _ := range hashMap {
		if !hashMap[num-1] {
			currentNum := num
			currentStreak := 1
			for hashMap[currentNum+1] {
				currentNum++
				currentStreak++
			}
			if max < currentStreak {
				max = currentStreak
			}
		}
	}
	return max
}
