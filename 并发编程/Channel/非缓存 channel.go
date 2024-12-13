package main

import "fmt"

func main() {
	done := make(chan struct{})

	go func() {
		fmt.Println("工作完成")
		done <- struct{}{}
	}()

	<-done // 等待信号
	fmt.Println("主 goroutine 继续执行")
}
