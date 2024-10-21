package main

import "fmt"

// 3 3 3
// .11
// .*.
// 22.
// 4

// 小红正在玩一个游戏。游戏的地图是一个n*m的迷宫，迷宫有墙和道路，道路上可能会有一些怪物。
// 小红初始的血量是，每当小红经过一个有怪物的道路时，小红就会和怪物战斗，击杀怪物并且消耗自己的血量。
// 小红消耗的血量等同于该怪物的战斗力。请注意，如果小红血量为0则死亡。因此只有当小红当前血量大于怪物的战斗力时才可经过该点。
// 地图共有以下几种标识： '.' 代表道路，小红可以经过。 '*' 代表墙体，小红不能经过。 '1'~'9' 数字，代表该位置是个道路，且上面有一个战斗力为该数字的怪物。
// 小红只可以上下左右四个方向移动。
// 小红想知道，自己从左上角到右下角的最短行走路线的距离是多少？

func main() {
	var n, m, h int
	fmt.Scan(&n, &m, &h)
	nums := make([][]int, n)
	for i := 0; i < n; i++ {
		nums[i] = make([]int, m)
	}
	var s string
	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		for j := 0; j < m; j++ {
			if s[j] == '.' {
				nums[i][j] = 0
			} else if s[j] == '*' {
				nums[i][j] = -1
			} else {
				nums[i][j] = int(s[j] - '0')
			}
		}
	}
	if nums[0][0] == -1 || nums[0][0] >= h {
		fmt.Println(-1)
		return
	}

	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dpM := make([][]int, m)
		for j := 0; j < m; j++ {
			dpM[j] = make([]int, 2)
		}
		dp[i] = dpM
	}
	dp[0][0][0] = 0
	dp[0][0][1] = nums[0][0]
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if nums[i][j] != -1 {
				if i > 0 && i-1 > 0 && i+1 <= n-1 {

				}

			}
		}
	}
	fmt.Println(dp[n-1][m-1][0])
}

func overStepCheck(x, y int) bool {

}
