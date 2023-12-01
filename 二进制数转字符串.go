package main

import "fmt"

func main() {
	fmt.Println(printBin(0.72))
}
func printBin(num float64) string {
	if num <= 0 || num >= 1 {
		return "ERROR"
	}
	binary := "0."
	for num > 0 {
		// 设置最大精度为32位
		if len(binary) >= 32 {
			return "ERROR"
		}
		// 将小数部分乘以2
		num *= 2
		if num >= 1 {
			// 如果乘积大于等于1，则将1添加到二进制表示中
			binary += "1"
			num -= 1
		} else {
			// 如果乘积小于1，则将0添加到二进制表示中
			binary += "0"
		}
	}

	return binary
}
