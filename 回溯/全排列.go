package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
func permute(nums []int) [][]int {
	var res [][]int
	var backtracking func(numsCopy []int, path []int, use []int)
	use := make([]int, len(nums))
	backtracking = func(numsCopy []int, path []int, use []int) {
		if len(path) == len(numsCopy) {
			temp := make([]int, len(numsCopy))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := 0; i < len(numsCopy); i++ {
			// 取过的就不取
			if use[i] == 0 {
				path = append(path, numsCopy[i])
				use[i] = 1
				backtracking(numsCopy, path, use)
				path = path[:len(path)-1]
				use[i] = 0
			}
		}
	}
	backtracking(nums, []int{}, use)
	return res
}
