package main

import "fmt"

func main() {
	safeDivision(10, 0) // 触发恐慌
	fmt.Println("程序继续执行")
}
func safeDivision(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("恢复了恐慌:", r)
		}
	}()

	if b == 0 {
		panic("除以零错误") // 引发恐慌
	}

	fmt.Println("结果:", a/b)
}
