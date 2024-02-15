package main

import "fmt"

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	if n == 0 {
		return []int{}
	}
	stack := []int{0}
	result := make([]int, n)
	for i := 1; i < n; i++ {
		top := stack[len(stack)-1]
		if temperatures[i] <= temperatures[top] {
			stack = append(stack, i)
		} else {
			// len(stack) > 0 防止死循环
			for len(stack) > 0 && temperatures[i] > temperatures[top] {
				result[top] = i - top
				stack = stack[:len(stack)-1]
				if len(stack) > 0 {
					top = stack[len(stack)-1]
				}
			}
			stack = append(stack, i)

		}
	}
	return result
}
