package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
}
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0, len(intervals))
	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if right < intervals[i][0] {
			res = append(res, []int{left, right})
			left, right = intervals[i][0], intervals[i][1]
		} else {
			right = max(right, intervals[i][1])
		}
	}
	res = append(res, []int{left, right}) // 将最后一个区间放入
	return res
}
