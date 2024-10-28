package main

//给你一个长度为 n 的整数数组 nums ，请你求出每个长度为 k 的子数组的 美丽值 。
//
//一个子数组的 美丽值 定义为：如果子数组中第 x 小整数 是 负数 ，那么美丽值为第 x 小的数，否则美丽值为 0 。
//
//请你返回一个包含 n - k + 1 个整数的数组，依次 表示数组中从第一个下标开始，每个长度为 k 的子数组的 美丽值 。
//
//    子数组指的是数组中一段连续 非空 的元素序列。
// 输入：nums = [1,-1,-3,-2,3], k = 3, x = 2
//输出：[-1,-2,-2]
//解释：总共有 3 个 k = 3 的子数组。
//第一个子数组是 [1, -1, -3] ，第二小的数是负数 -1 。
//第二个子数组是 [-1, -3, -2] ，第二小的数是负数 -2 。
//第三个子数组是 [-3, -2, 3] ，第二小的数是负数 -2 。
import "fmt"

func main() {
	fmt.Println(getSubarrayBeauty([]int{-3, 1, 2, -3, 0, -3}, 2, 1))
}
func getSubarrayBeauty(nums []int, k int, x int) []int {
	bias := 50
	win := make([]int, bias*2+1)
	ans := make([]int, len(nums)-k+1)
	for i := 0; i < k-1; i++ {
		win[nums[i]+bias]++
	}
	for i := k - 1; i < len(nums); i++ {
		win[nums[i]+bias]++
		left := x
		for index, v := range win {
			left -= v
			if left <= 0 {
				if index-bias < 0 {
					ans[i-k+1] = index - bias
				}
				break
			}
		}
		win[nums[i-k+1]+bias]--
	}
	return ans
}
