package main

import "fmt"

// 给你一个由 n 个元素组成的整数数组 nums 和一个整数 k 。
// 请你找出平均数最大且 长度为 k 的连续子数组，并输出该最大平均数。
// 任何误差小于 10-5 的答案都将被视为正确答案。
func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
}

//输入：nums = [1,12,-5,-6,50,3], k = 4
//输出：12.75
//解释：最大平均数 (12-5-6+50)/4 = 51/4 = 12.75

// 输入：nums = [5], k = 1
// 输出：5.00000
func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	var res float64
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	res = float64(sum) / float64(k)
	for i := k; i < len(nums); i++ {
		sum -= nums[i-k]
		sum += nums[i]
		if res < float64(sum)/float64(k) {
			res = float64(sum) / float64(k)
		}
		if i+1 < k {
			continue
		}
	}
	return res
}
