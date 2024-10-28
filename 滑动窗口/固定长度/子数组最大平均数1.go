package main

import "fmt"

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
}
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
