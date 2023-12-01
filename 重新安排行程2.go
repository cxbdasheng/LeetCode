package main

import (
	"fmt"
	"sort"
)

func main() {
	tickets := [][]string{
		{"JFK", "KUL"},
		{"JFK", "NRT"},
		{"NRT", "JFK"},
	}
	sort.Sort(ticketsSli(tickets))
	fmt.Println(tickets)
	use := make([]int, len(tickets))
	path = append(path, "JFK")
	fmt.Println(backtracking(tickets, "JFK", use))
}

// var result []string
var path []string

func backtracking(tickets [][]string, start string, use []int) []string {

	for index, value := range tickets {
		if len(path) == len(tickets)+1 {
			return path
		}
		if use[index] == 1 {
			continue
		}
		if value[0] == start {
			use[index] = 1
			if len(path) >= 1 && path[len(path)-1] != start {
				path = append(path, start)
			}
			path = append(path, value[1])
			backtracking(tickets, value[1], use)
		}
	}
	return path
}

type ticketsSli [][]string

func (this ticketsSli) Less(i, j int) bool {
	if this[i][0] == this[j][0] {

		return this.isSort(this[i][1], this[j][1])
	}
	return this.isSort(this[i][0], this[j][0])
}
func (this ticketsSli) Len() int {
	return len(this)
}
func (this ticketsSli) Swap(i, j int) {
	(this)[i], (this)[j] = (this)[j], (this)[i]
}
func (this ticketsSli) isSort(a string, b string) bool {
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
