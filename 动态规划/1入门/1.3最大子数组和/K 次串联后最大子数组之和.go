package main

import "fmt"

// 给定一个整数数组 arr 和一个整数 k ，通过重复 k 次来修改数组。
// 例如，如果 arr = [1, 2] ， k = 3 ，那么修改后的数组将是 [1, 2, 1, 2, 1, 2] 。
// 返回修改后的数组中的最大的子数组之和。注意，子数组长度可以是 0，在这种情况下它的总和也是 0。
// 由于 结果可能会很大，需要返回的 109 + 7 的 模 。
func main() {
	fmt.Println(kConcatenationMaxSum([]int{1, 2}, 1))
}

// 输入：arr = [1,2], k = 3
// 输出：9

// 输入：arr = [1,-2,1], k = 5
// 输出：2

// 输入：arr = [-1,-2], k = 7
// 输出：0

func kConcatenationMaxSum(arr []int, k int) int {
	const mod = 1_000_000_007
	var ans int
	r := 0
	s := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		r = max(r, 0) + arr[i]
		ans = max(ans, r)
		s = (s + arr[i]) % mod
	}
	if k == 1 {
		return ans % mod
	}
	for i := n; i < n*2; i++ {
		r = max(r, 0) + arr[i%n]
		ans = max(ans, r)
	}
	if k == 2 {
		return ans % mod
	}
	if s > 0 {
		return (ans%mod + ((k-2)*s)%mod) % mod
	}
	return ans
}
