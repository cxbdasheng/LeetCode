package main

import "fmt"

func main() {
	fmt.Println(getAverages([]int{7, 4, 3, 9, 1, 8, 5, 2, 6}, 3))
}
func getAverages(nums []int, k int) []int {
	//res := make([]int, len(nums))
	//for i := 0; i < len(nums); i++ {
	//	res[i] = -1
	//}
	//if len(nums) < k*2+1 {
	//	return res
	//}
	//sum := 0
	//for i := 0; i < 2*k; i++ {
	//	sum += nums[i]
	//}
	//for i := k; i < len(nums); i++ {
	//	if i-k >= 0 && i+k <= len(nums)-1 {
	//		sum += nums[i+k]
	//		res[i] = sum / (2*k + 1)
	//		sum -= nums[i-k]
	//	}
	//}
	//return res
	avgs := make([]int, len(nums))
	sum := 0
	for i, v := range nums {
		if i < k || i+k >= len(nums) { // 超过边界
			avgs[i] = -1
		}
		sum += v
		if i >= k*2 {
			avgs[i-k] = sum / (k*2 + 1)
			sum -= nums[i-k*2]
		}
	}
	return avgs
}
