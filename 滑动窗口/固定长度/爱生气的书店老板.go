package main

import "fmt"

func main() {
	fmt.Println(maxSatisfied([]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3))
}
func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	m := 0
	sum := 0
	ans := 0
	for i := 0; i < minutes; i++ {
		if grumpy[i] == 1 {
			sum += customers[i]
		} else {
			ans += customers[i]
		}
	}
	m = sum
	for i := minutes; i < len(customers); i++ {
		if grumpy[i] == 1 {
			sum += customers[i]
		} else {
			ans += customers[i]
		}
		if grumpy[i-minutes] == 1 {
			sum -= customers[i-minutes]
		}
		if sum > m {
			m = sum
		}
	}
	return ans + m
}
