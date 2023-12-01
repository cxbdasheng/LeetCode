package main

import (
	"fmt"
)

func main() {
	tickets := [][]string{
		{"MUC", "LHR"},
		{"JFK", "MUC"},
		{"SFO", "SJC"},
		{"LHR", "SFO"},
	}
	result = []string{}
	var start = "JFK"
	result = append(result, start)
	use := make([]int, len(tickets))
	fmt.Println(backtracking(tickets, start, use))
}

var result []string

func backtracking(tickets [][]string, start string, use []int) []string {
	to := make(map[int]string)
	min := 0
	for index, value := range tickets {
		if use[index] == 1 {
			continue
		}
		if value[0] == start {
			to[index] = value[1]
			min = index
		}
	}
	if len(to) == 0 {
		return result
	}
	if len(to) > 1 {
		min = toMin(to, use)
	}
	use[min] = 1
	result = append(result, to[min])
	backtracking(tickets, to[min], use)
	return result
}

func toMin(to map[int]string, use []int) int {
	min := -1
	for index, value := range to {
		if use[index] == 1 {
			continue
		}
		if min == -1 {
			min = index
			continue
		}
		if isSort(value, to[min]) {
			min = index
		}
	}
	return min
}

/*
*
排序靠前的为true
*/
func isSort(a string, b string) bool {
	var n = len(b)
	if len(a) > len(b) {
		n = len(a)
	}
	for i := 0; i < n; i++ {
		if a[i] < b[i] {
			return true
		} else if a[i] > b[i] {
			return false
		}
	}

	return true
}
