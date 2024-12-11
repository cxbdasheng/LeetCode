package main

import "fmt"

func main() {
	outerFunc()
	//printNumbers1()
	//fmt.Println()
	//printNumbers2()
	printNumbers3()
}

func outerFunc() {
	defer fmt.Println("函数执行结束前一刻才会被打印。")
	fmt.Println("第一个被打印的。")
}

// 外围函数中执行顺序不同
func printNumbers1() {
	for i := 0; i < 10; i++ {
		defer func(n int) {
			fmt.Printf("%d", n)
		}(i)
	}
}

// 多个延迟函数调用的执行顺序不同
func printNumbers2() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
		defer fmt.Println(i * 2)
	}
}

// 语句执行时求出 i*2
func printNumbers3() {
	for i := 0; i < 5; i++ {
		defer func(n int) {
			fmt.Printf("%d \n", n)
		}(i * 2)
	}
}
