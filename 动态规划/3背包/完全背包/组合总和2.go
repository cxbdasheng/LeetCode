package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
}
func combinationSum4(nums []int, target int) int {
	input := bufio.NewScanner(os.Stdin)
	dp := make([]int, target+1)
	dp[0] = 1
	for j := 0; j <= target; j++ {
		for i := 0; i < len(nums); i++ {
			if nums[i] <= j {
				dp[j] = dp[j] + dp[j-nums[i]]
			}
		}
	}
	return dp[target]
}
