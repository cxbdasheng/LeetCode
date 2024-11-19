package main

import "fmt"

//给定一个整数数组  nums，处理以下类型的多个查询:
//计算索引 left 和 right （包含 left 和 right）之间的 nums 元素的 和 ，其中 left <= right
//实现 NumArray 类：
//NumArray(int[] nums) 使用数组 nums 初始化对象
//int sumRange(int i, int j) 返回数组 nums 中索引 left 和 right 之间的元素的 总和 ，包含 left 和 right 两点（也就是 nums[left] + nums[left + 1] + ... + nums[right] )

func main() {
	obj := Constructor([]int{1, 2, 3, 4, 5})
	fmt.Println(obj.SumRange(0, 6))
}

// 输入：
// ["NumArray", "sumRange", "sumRange", "sumRange"]
// [[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
// 输出：
// [null, 1, -1, -3]
type NumArray []int

func Constructor(nums []int) NumArray {
	// 填充 0 不需要特殊处理
	s := make(NumArray, len(nums)+1)
	for i, x := range nums {
		s[i+1] = s[i] + x
	}
	return s
}

func (s NumArray) SumRange(left, right int) int {
	if left > right {
		return 0
	}
	if left < 0 || right >= len(s) {
		return -1
	}
	return s[right+1] - s[left]
}
