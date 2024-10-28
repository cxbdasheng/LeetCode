package main

import "fmt"

//几张卡牌 排成一行，每张卡牌都有一个对应的点数。点数由整数数组 cardPoints 给出。
//
//每次行动，你可以从行的开头或者末尾拿一张卡牌，最终你必须正好拿 k 张卡牌。
//
//你的点数就是你拿到手中的所有卡牌的点数之和。
//
//给你一个整数数组 cardPoints 和整数 k，请你返回可以获得的最大点数。
//输入：cardPoints = [1,2,3,4,5,6,1], k = 3
//输出：12
// 逆向思维：n-k 滑动窗口的最大值

func main() {
	fmt.Println(maxScore([]int{100, 40, 17, 9, 73, 75}, 3))
}
func maxScore(cardPoints []int, k int) int {
	m := len(cardPoints) - k
	sum := 0
	total := 0
	for i := 0; i < m; i++ {
		sum += cardPoints[i]
		total += cardPoints[i]
	}
	low := sum
	for i := m; i < len(cardPoints); i++ {
		sum += cardPoints[i]
		sum -= cardPoints[i-m]
		if sum < low {
			low = sum
		}
		total += cardPoints[i]
	}
	return total - low
}
