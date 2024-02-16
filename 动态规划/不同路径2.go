package main

import "fmt"

func main() {
	var obstacleGrid = [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	fmt.Println(uniquePathsWithObstacles(obstacleGrid))
}
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) < 0 {
		return 0
	}
	dp := make([][]int, len(obstacleGrid))
	for i := 0; i < len(obstacleGrid); i++ {
		dp[i] = make([]int, len(obstacleGrid[i]))
	}
	dp[0][0] = 1
	if obstacleGrid[0][0] == 1 {
		dp[0][0] = 0
	}
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[i]); j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			if i > 0 {
				dp[i][j] += dp[i-1][j] //当前点左侧有点
			}
			if j > 0 {
				dp[i][j] += dp[i][j-1] //同上，上方的点
			}
		}
	}
	return dp[len(obstacleGrid)-1][len(obstacleGrid[len(obstacleGrid)-1])-1]
}
