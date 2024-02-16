package main

import (
	"fmt"
	"strings"
)

func main() {
	var n = 4
	chessboard := make([][]string, n)
	for i := 0; i < n; i++ {
		chessboard[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			chessboard[i][j] = "."
		}
	}
	fmt.Println(backtracking(0, n, chessboard))
}

var result [][]string

func backtracking(row int, n int, chessboard [][]string) [][]string {
	if row == n {
		temp := make([]string, n)
		for i, rowStr := range chessboard {
			temp[i] = strings.Join(rowStr, "")
		}
		result = append(result, temp)
		return result
	}
	for i := 0; i < n; i++ {
		if row == 0 || isValid(row, i, chessboard, n) {
			chessboard[row][i] = "Q"
			backtracking(row+1, n, chessboard)
			chessboard[row][i] = "."
		}
	}
	return result
}

func isValid(row int, col int, chessboard [][]string, n int) bool {
	for i := 0; i < row; i++ {
		if chessboard[i][col] == "Q" {
			return false
		}
	}
	// 45度
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}
	// 135度
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}
	return true

}
