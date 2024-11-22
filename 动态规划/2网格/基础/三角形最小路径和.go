package main

import (
	"fmt"
	"slices"
)

// 给定一个三角形 triangle ，找出自顶向下的最小路径和。
// 每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。
// 也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。
func main() {
	/**
	递推
	*/
	fmt.Println(minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
	/**
	BFS
	*/
	fmt.Println(minimumTotal1([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))

}

// 输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
// 输出：11
// 解释：如下面简图所示：
//
//	  2
//	 3 4
//	6 5 7
//
// 4 1 8 3
// 自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。

// 输入：triangle = [[-10]]
// 输出：-10
func minimumTotal(triangle [][]int) int {
	dp := make([][]int, len(triangle))
	for i := range dp {
		dp[i] = make([]int, len(triangle[i]))
	}
	for i := 0; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if i == 0 && j == 0 {
				dp[i][j] = triangle[i][j]
				continue
			}
			if j-1 >= 0 && j < len(triangle[i-1]) {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j == len(triangle[i])-1 {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			}
		}
	}
	fmt.Println(dp)
	return slices.Min(dp[len(dp)-1])
}

/*
*
DFS
*/
func minimumTotal1(triangle [][]int) int {
	var dfs func(int, int) int
	flag := make([][]int, len(triangle))
	memo := make([][]int, len(triangle))
	for i := range memo {
		memo[i] = make([]int, len(triangle[i]))
		flag[i] = make([]int, len(triangle[i]))
	}
	dfs = func(i, j int) int {
		if flag[i][j] != 0 {
			return memo[i][j]
		}
		if i == len(triangle)-1 {
			return triangle[i][j]
		}
		memo[i][j] = min(dfs(i+1, j)+triangle[i][j], dfs(i+1, j+1)+triangle[i][j])
		flag[i][j] = 1
		return memo[i][j]
	}
	return dfs(0, 0)
}
