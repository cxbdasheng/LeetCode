package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	var parts []string
	if input.Scan() {
		parts = strings.Split(input.Text(), " ")
		if len(parts) < 2 {
			return
		}
	}
	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])
	dp := make([]int, n+1)
	dp[0] = 1
	for j := 1; j <= n; j++ {
		for i := 1; i <= m; i++ {
			if i <= j {
				dp[j] = dp[j] + dp[j-i]
			}
		}
	}
	fmt.Println(dp[n])
}
