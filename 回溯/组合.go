package main

import "fmt"

func main() {
	fmt.Println(combine(16, 16, 1))
}

var path []int
var result [][]int

func combine(n int, k int, index int) [][]int {
	if sumArray(path) == k {
		temp := make([]int, k)
		copy(temp, path)
		result = append(result, temp)
		return result
	}
	for i := index; i <= n; i++ {
		path = append(path, i)
		combine(n, k, i+1)
		path = path[:len(path)-1]
	}
	return result
}

// 求和函数
func sumArray(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}
