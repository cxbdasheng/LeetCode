package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findContentChildren([]int{1, 2}, []int{1, 2, 3}))
}
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var sum int
	r := len(g) - 1
	for i := len(s) - 1; i >= 0; i-- {
		if r < 0 {
			break
		}
		for j := r; j >= 0; j-- {
			if s[i] >= g[j] {
				sum++
				r, j = j-1, -1
			}
		}
	}
	return sum
}
