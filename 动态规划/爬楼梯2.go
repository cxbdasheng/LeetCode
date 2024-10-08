package main

import "fmt"

func main() {
	//input := bufio.NewScanner(os.Stdin)
	//var parts []string
	//if input.Scan() {
	//	parts = strings.Split(input.Text(), " ")
	//	if len(parts) < 2 {
	//		return
	//	}
	//}
	//n, _ := strconv.Atoi(parts[0])
	//m, _ := strconv.Atoi(parts[1])
	//dp := make([]int, n+1)
	//dp[0] = 1
	//for j := 1; j <= n; j++ {
	//	for i := 1; i <= m; i++ {
	//		if i <= j {
	//			dp[j] = dp[j] + dp[j-i]
	//		}
	//	}
	//}
	//fmt.Println(dp[n])
	fmt.Println(minCostClimbingStairs([]int{1, 100}))
}
func minCostClimbingStairs(cost []int) int {
	if len(cost) == 0 || len(cost) == 1 {
		return 0
	}
	dp := make([]int, len(cost)+1)
	dp[0] = 0
	dp[1] = 0
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[len(cost)]
}
