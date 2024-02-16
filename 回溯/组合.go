package main

import "fmt"

func main() {
	fmt.Println(combine(4, 2, 1))
}

var path []int
var result [][]int

func combine(n int, k int, index int) [][]int {
	if len(path) == k {
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
