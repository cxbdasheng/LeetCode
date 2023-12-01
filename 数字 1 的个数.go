package main

import "fmt"

// 给定一个整数 num，计算所有小于等于 num 的非负整数中数字 1 出现的个数。
func main() {
	var num = 13
	var base, res int64 = 1, 0
	n64 := int64(num)
	for base <= n64 {
		// 计算 a..., cur, b...
		a := n64 / base / 10
		cur := (n64 / base) % 10
		b := n64 % base
		// 将当前位设为1，考察其他部分的变化范围
		if cur > 1 {
			// 一、cur > 1，
			//          [3101 ] 5 [92]
			// 变化范围：[0-3101] 1 [0-99]
			// 总个数：   (a+1)  *  base
			res += (a + 1) * base
		} else if cur == 1 {
			// 二、cur == 1，
			//             [310] 1 [592]
			// 1、变化范围 [0-309] 1 [0-999]
			//              a    *  base
			// 2、变化范围 [310]   1 [0-592]
			//               1   *   (b+1)
			// 总个数：a *base + (b + 1)
			res += a*base + b + 1

		} else {
			// 三、cur < 1，
			//           [31] 0 [1592]
			// 变化范围 [0-30] 1 [0-9999]
			// 总个数    a     *   base
			res += a * base
		}
		// 统计更高一位
		base *= 10
	}
	fmt.Println(res)
}
