package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 1)
	ticker := time.NewTicker(time.Second)
	go func() {
		for _ = range ticker.C {
			select {
			case intChan <- 1:
			case intChan <- 2:
			case intChan <- 3:
			}
		}
		close(intChan)
		fmt.Println("结束发送")
	}()

	var sum int

	for v := range intChan {
		fmt.Printf("读取到值：%v \n", v)
		sum += v
		if sum > 10 {
			fmt.Println("大于 10 结束")
			break
		}
	}
	fmt.Println("结束接收")
}
