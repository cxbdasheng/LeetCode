package main

import "fmt"

func main() {
	var chessboard = [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	backtracking(chessboard)
	fmt.Println(chessboard)
}

func backtracking(chessboard [][]byte) bool {

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if chessboard[i][j] != '.' {
				continue
			}
			for k := '1'; k <= '9'; k++ {
				if isvalid(i, j, byte(k), chessboard) {
					chessboard[i][j] = byte(k)
					if backtracking(chessboard) == true {
						return true
					}
					chessboard[i][j] = '.'
				}
			}
			return false
		}
	}
	return true
}

// [53 51 52 54 55 56 57 49 50] [54 55 50 49 57 53 51 52 56] [49 57 56 51 52 50 53 54 55] [56 53 57 55 54 49 52 50 51] [52 50 54 56 53 51 55 57 49] [55 49 51 57 50 52 56 53 54] [57 54 49 53 51 55 50 56 52] [50 56 55 52 49 57 54 51 53] [51 52 53 50 56 54 49 55 57]
//func isValid(row int, col int, k byte, chessboard [][]byte) bool {
//	// 行
//	for _, v := range chessboard[row] {
//		if v == k {
//			return false
//		}
//	}
//
//	// 列
//	for i := 0; i < len(chessboard); i++ {
//		if chessboard[i][col] == k {
//			return false
//		}
//	}
//	// 3*3
//	r := row / 3
//	c := col / 3
//	for i := r * 3; i < r+3; i++ {
//		for j := c * 3; j < c+3; j++ {
//			if chessboard[i][j] == k {
//				return false
//			}
//		}
//	}
//
//	return true
//}

func isvalid(row, col int, k byte, board [][]byte) bool {
	for i := 0; i < 9; i++ { //行
		if board[row][i] == k {
			return false
		}
	}
	for i := 0; i < 9; i++ { //列
		if board[i][col] == k {
			return false
		}
	}
	//方格
	startrow := (row / 3) * 3
	startcol := (col / 3) * 3
	for i := startrow; i < startrow+3; i++ {
		for j := startcol; j < startcol+3; j++ {
			if board[i][j] == k {
				return false
			}
		}
	}
	return true
}
