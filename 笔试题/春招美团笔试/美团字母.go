package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	var s string
	if input.Scan() {
		s = input.Text()
	}
	var l int
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			l++
		}
	}
	var result int
	if len(s)-l < l {
		result = len(s) - l
	} else {
		result = l
	}
	if s[0] >= 'A' && s[0] <= 'Z' {
		if l-1 < result {
			result = l - 1
		}
	}
	fmt.Println(result)

	//l := 0
	//"AbbC"
	//
	//for input.Scan() {
	//	s = input.Text()
	//	dp := make([][]int, len(s))
	//	for i := 0; i < len(s); i++ {
	//		dp[i] = make([]int, 3)
	//	}
	//
	//	if s[0] >= 'a' && s[0] <= 'z' {
	//		dp[0][0] = 1
	//		dp[0][1] = 0
	//		dp[0][2] = 1
	//	} else {
	//		dp[0][0] = 0
	//		dp[0][1] = 1
	//		dp[0][2] = 0
	//	}
	//
	//	for i := 1; i < len(s); i++ {
	//		if s[i] >= 'a' && s[i] <= 'Z' {
	//			dp[i][0] = dp[i-1][0]
	//			dp[i][1] = dp[i-1][1]
	//			dp[i][2] = dp[i-1][2] + 1
	//		} else {
	//			dp[i][0] = dp[i-1][0] + 1
	//			dp[i][1] = dp[i-1][1] + 1
	//			dp[i][2] = dp[i-1][2]
	//		}
	//	}
	//	for i := 0; i < len(dp); i++ {
	//		fmt.Println(dp[i])
	//	}
	//	if dp[len(s)-1][0] < dp[len(s)-1][1] {
	//		if dp[len(s)-1][0] <= dp[len(s)-1][2] {
	//			fmt.Println(dp[len(s)-1][0])
	//			continue
	//		}
	//	}
	//	if dp[len(s)-1][1] < dp[len(s)-1][0] {
	//		if dp[len(s)-1][1] <= dp[len(s)-1][2] {
	//			fmt.Println(dp[len(s)-1][1])
	//			continue
	//		}
	//	}
	//	if dp[len(s)-1][2] < dp[len(s)-1][0] {
	//		if dp[len(s)-1][2] <= dp[len(s)-1][1] {
	//			fmt.Println(dp[len(s)-1][2])
	//			continue
	//		}
	//	}
	//}
}
