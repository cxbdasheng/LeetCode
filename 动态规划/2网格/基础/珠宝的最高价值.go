package main

import "fmt"

// 现有一个记作二维矩阵 frame 的珠宝架，其中 frame[i][j] 为该位置珠宝的价值。拿取珠宝的规则为：
// 只能从架子的左上角开始拿珠宝
// 每次可以移动到右侧或下侧的相邻位置
// 到达珠宝架子的右下角时，停止拿取
// 注意：珠宝的价值都是大于 0 的。除非这个架子上没有任何珠宝，比如 frame = [[0]]。
func main() {
	fmt.Println(jewelleryValue([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}
func jewelleryValue(frame [][]int) int {
	dp := make([][]int, len(frame))
	for i := range dp {
		dp[i] = make([]int, len(frame[i]))
	}
	for i := 0; i < len(frame); i++ {
		for j := 0; j < len(frame[i]); j++ {
			if i == 0 && j == 0 {
				dp[i][j] = frame[i][j]
				continue
			}
			if i == 0 {
				dp[i][j] = dp[i][j-1] + frame[i][j]
				continue
			}
			if j == 0 {
				dp[i][j] = dp[i-1][j] + frame[i][j]
				continue
			}
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + frame[i][j]
		}
	}
	return dp[len(frame)-1][len(frame[0])-1]
}
