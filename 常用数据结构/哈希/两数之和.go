package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, num := range nums {
		if p, ok := mp[target-num]; ok {
			return []int{p, i}
		}
		mp[num] = i
	}
	return nil
}
