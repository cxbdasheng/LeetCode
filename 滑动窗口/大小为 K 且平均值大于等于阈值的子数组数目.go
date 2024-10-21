package main

import "fmt"

func main() {
	fmt.Println(numOfSubarrays([]int{1, 1, 1, 1, 1}, 1, 0))
}
func numOfSubarrays(arr []int, k int, threshold int) int {
	sum := 0
	for i := 0; i < k; i++ {
		sum += arr[i]
	}
	res := 0
	if sum/k >= threshold {
		res++
	}
	for i := k; i < len(arr); i++ {
		sum -= arr[i-k]
		sum += arr[i]
		if sum/k >= threshold {
			res++
		}
	}
	return res
}
