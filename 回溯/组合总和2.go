package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum2([]int{14, 6, 25, 9, 30, 20, 33, 34, 28, 30, 16, 12, 31, 9, 9, 12, 34, 16, 25, 32, 8, 7, 30, 12, 33, 20, 21, 29, 24, 17, 27, 34, 11, 17, 30, 6, 32, 21, 27, 17, 16, 8, 24, 12, 12, 28, 11, 33, 10, 32, 22, 13, 34, 18, 12}, 27))
}
func combinationSum2(candidates []int, target int) [][]int {
	var ans [][]int
	var path []int
	sort.Ints(candidates)
	use := make([]bool, len(candidates))
	var backtrack func(int, int, []bool)
	backtrack = func(target int, start int, use []bool) {
		if target == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		if target < 0 {
			return
		}
		for i := start; i < len(candidates); i++ {
			if i > 0 && candidates[i] == candidates[i-1] && use[i-1] == false {
				continue
			}
			if candidates[i] > target {
				return
			}
			use[i] = true
			path = append(path, candidates[i])
			backtrack(target-candidates[i], i+1, use)
			path = path[:len(path)-1]
			use[i] = false
		}
		return
	}
	backtrack(target, 0, use)
	return ans
}
