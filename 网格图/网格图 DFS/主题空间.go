package main

import "fmt"

// 「以扣会友」线下活动所在场地由若干主题空间与走廊组成，场地的地图记作由一维字符串型数组 grid，字符串中仅包含 "0"～"5" 这 6 个字符。
// 地图上每一个字符代表面积为 1 的区域，其中 "0" 表示走廊，其他字符表示主题空间。相同且连续（连续指上、下、左、右四个方向连接）的字符组成同一个主题空间。
// 假如整个 grid 区域的外侧均为走廊。请问，不与走廊直接相邻的主题空间的最大面积是多少？如果不存在这样的空间请返回 0。
func main() {
	//fmt.Println(largestArea([]string{"110", "231", "221"}))
	fmt.Println(largestArea1([]string{"110", "231", "221"}))
}

// 输入：grid = ["110","231","221"]
// 输出：1
// 解释：4 个主题空间中，只有 1 个不与走廊相邻，面积为 1。

// 输入：grid = ["11111100000","21243101111","21224101221","11111101111"]
// 输出：3
// 解释：8 个主题空间中，有 5 个不与走廊相邻，面积分别为 3、1、1、1、2，最大面积为 3。
var count int

func largestArea(grid []string) int {

	var ans int
	arr := make([][]uint8, len(grid))
	vis := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		var charArr []uint8
		for j := 0; j < len(grid[i]); j++ {
			charArr = append(charArr, grid[i][j]-'0')
		}
		arr[i] = charArr
		vis[i] = make([]bool, len(grid[i]))
	}
	var dfs func(x, y int, v uint8, vis [][]bool) bool
	dfs = func(x, y int, v uint8, vis [][]bool) bool {
		if x < 0 || x >= len(arr) || y < 0 || y >= len(arr[x]) || arr[x][y] == 0 {
			return false
		}
		if arr[x][y] != v || vis[x][y] {
			return true
		}
		c := true
		vis[x][y] = true
		count++
		c = dfs(x+1, y, v, vis) && c
		c = dfs(x-1, y, v, vis) && c
		c = dfs(x, y+1, v, vis) && c
		c = dfs(x, y-1, v, vis) && c
		return c
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			count = 0
			if dfs(i, j, arr[i][j], vis) {
				ans = max(ans, count)
			}
		}
	}
	return ans
}

func largestArea1(grid []string) int {
	m, n := len(grid), len(grid[0])
	vis := [505][505]bool{}
	dx := [4]int{-1, 0, 1, 0}
	dy := [4]int{0, -1, 0, 1}
	var dfs func(v byte, x, y int) (int, bool)

	dfs = func(v byte, x, y int) (int, bool) {
		ret, flag := 1, true
		vis[x][y] = true
		for k := 0; k < 4; k++ {
			nx, ny := x+dx[k], y+dy[k]
			if nx < 0 || nx >= m || ny < 0 || ny >= n || grid[nx][ny] == '0' {
				flag = false
				continue
			}
			if vis[nx][ny] || grid[nx][ny] != v {
				continue
			}
			s, t := dfs(v, nx, ny)
			if !t {
				flag = false
			}
			ret += s
		}
		return ret, flag
	}

	ans := 0
	for i, r := range grid {
		for j, v := range r {
			if vis[i][j] {
				continue
			}
			if s, flag := dfs(byte(v), i, j); flag {
				ans = max(ans, s)
			}
		}
	}
	return ans
}
