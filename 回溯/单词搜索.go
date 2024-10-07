package main

import "fmt"

func main() {
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB"))
}

var dirs = []struct{ x, y int }{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func exist(board [][]byte, word string) bool {

	var backtrack func(i, j, index int) bool
	backtrack = func(i, j, index int) bool {
		if board[i][j] != word[index] {
			return false
		}
		if index == len(word)-1 {
			return true
		}
		board[i][j] = 0
		for _, d := range dirs {
			x, y := i+d.x, j+d.y // 相邻格子
			if 0 <= x && x < len(board) && 0 <= y && y < len(board[x]) && backtrack(x, y, index+1) {
				return true // 搜到了！
			}
		}
		board[i][j] = word[index] // 恢复现场
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if backtrack(i, j, 0) {
				return true
			}
		}
	}
	return false
}
