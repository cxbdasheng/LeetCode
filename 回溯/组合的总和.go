package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var path []int
	sort.Ints(candidates)
	var backtracking func(candidates []int, target int, start int)
	backtracking = func(candidates []int, target int, start int) {
		if target == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := start; i < len(candidates); i++ {
			if candidates[i] <= target {
				path = append(path, candidates[i])
				target = target - candidates[i]
				backtracking(candidates, target, i)
				target = target + candidates[i]
				path = path[:len(path)-1]
			} else {
				break
			}
		}
	}
	backtracking(candidates, target, 0)
	return res
}
