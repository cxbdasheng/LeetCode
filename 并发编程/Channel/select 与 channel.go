package main

import (
	"fmt"
)

func main() {
	//ch1 := make(chan int, 3)
	//ch2 := make(chan int, 3)
	//close(ch1)
	//select {
	//case elem := <-ch1:
	//	fmt.Printf("ch1 通道值%v", elem)
	//case elem := <-ch2:
	//	fmt.Printf("ch2 通道值%v", elem)
	//default:
	//	fmt.Println("没有通道数据")
	//}

	ch1 := make(chan string, 5)
	for i := 0; i < 5; i++ {
		ch1 <- fmt.Sprintf("来自 ch1: %d", i)
	}
	close(ch1)
	syncChan := make(chan struct{}, 1)
	// 使用 select 处理多个通道
	go func() {
	Loop:
		for {
			select {
			case msg1, ok := <-ch1:
				if !ok {
					fmt.Println("结束了 for Loop 循环")
					break Loop
				}
				fmt.Printf("%s \n", msg1)
			default:
				fmt.Printf("运行了默认逻辑")
			}

		}
		syncChan <- struct{}{}
	}()
	<-syncChan
	fmt.Println("所有消息已接收，结束程序。")
}
