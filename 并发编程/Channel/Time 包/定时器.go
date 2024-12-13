package main

import (
	"fmt"
	"time"
)

func main() {
	//t := time.NewTimer(time.Second * 3)
	//fmt.Printf("当前时间：%v \n", time.Now())
	//exp := <-t.C
	//fmt.Printf("失效时间：%v \n", exp)
	//fmt.Printf("结束时间：%v \n", t.Stop())

	//intChan := make(chan int, 1)
	//go func() {
	//	time.Sleep(time.Second)
	//	intChan <- 1
	//}()
	//
	//select {
	//case i := <-intChan:
	//	fmt.Printf("接收到元素的值：%v\n", i)
	//case <-time.NewTimer(time.Millisecond * 500).C:
	//	fmt.Println("已经超时")
	//}

	intChan := make(chan int, 1)
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			intChan <- i
		}
		close(intChan)
	}()
	timeout := time.Millisecond * 500
	var timer *time.Timer
	for {
		if timer == nil {
			timer = time.NewTimer(timeout)
		} else {
			timer.Reset(timeout)
		}
		select {
		case i, ok := <-intChan:
			if !ok {
				fmt.Println("结束")
				return
			}
			fmt.Printf("接收到元素的值：%v\n", i)
		case <-timer.C:
			fmt.Println("已经超时")
		}
	}
}
