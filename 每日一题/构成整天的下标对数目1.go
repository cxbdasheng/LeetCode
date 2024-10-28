package main

import "fmt"

func main() {
	fmt.Println(countCompleteDayPairs([]int{12, 12, 30, 24, 24}))
}
func countCompleteDayPairs(hours []int) int {
	j := 0
	ans := 0
	for j < len(hours) {
		for i := j + 1; i < len(hours); i++ {
			if (hours[i]+hours[j])%24 == 0 {
				ans++
			}
		}
		j++
	}
	return ans
}
