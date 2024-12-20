package main

import "fmt"

//给你一个下标从 0 开始的数组 nums ，数组中有 n 个整数，另给你一个整数 k 。
//半径为 k 的子数组平均值 是指：nums 中一个以下标 i 为 中心 且 半径 为 k 的子数组中所有元素的平均值，即下标在 i - k 和 i + k 范围（含 i - k 和 i + k）内所有元素的平均值。如果在下标 i 前或后不足 k 个元素，那么 半径为 k 的子数组平均值 是 -1 。
//构建并返回一个长度为 n 的数组 avgs ，其中 avgs[i] 是以下标 i 为中心的子数组的 半径为 k 的子数组平均值 。
//x 个元素的 平均值 是 x 个元素相加之和除以 x ，此时使用截断式 整数除法 ，即需要去掉结果的小数部分。
//例如，四个元素 2、3、1 和 5 的平均值是 (2 + 3 + 1 + 5) / 4 = 11 / 4 = 2.75，截断后得到 2 。

func main() {
	fmt.Println(getAverages([]int{7, 4, 3, 9, 1, 8, 5, 2, 6}, 3))
}

// 输入：nums = [7,4,3,9,1,8,5,2,6], k = 3
// 输出：[-1,-1,-1,5,4,4,-1,-1,-1]
// 解释：
// - avg[0]、avg[1] 和 avg[2] 是 -1 ，因为在这几个下标前的元素数量都不足 k 个。
// - 中心为下标 3 且半径为 3 的子数组的元素总和是：7 + 4 + 3 + 9 + 1 + 8 + 5 = 37 。
// 使用截断式 整数除法，avg[3] = 37 / 7 = 5 。
// - 中心为下标 4 的子数组，avg[4] = (4 + 3 + 9 + 1 + 8 + 5 + 2) / 7 = 4 。
// - 中心为下标 5 的子数组，avg[5] = (3 + 9 + 1 + 8 + 5 + 2 + 6) / 7 = 4 。
// - avg[6]、avg[7] 和 avg[8] 是 -1 ，因为在这几个下标后的元素数量都不足 k 个。

// 输入：nums = [100000], k = 0
// 输出：[100000]
// 解释：
// - 中心为下标 0 且半径 0 的子数组的元素总和是：100000 。
// avg[0] = 100000 / 1 = 100000 。
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
		// 超过边界
		if i < k || i+k >= len(nums) {
			avgs[i] = -1
		}
		sum += v
		//
		if i >= k*2 {
			avgs[i-k] = sum / (k*2 + 1)
			sum -= nums[i-k*2]
		}
	}
	return avgs
}
