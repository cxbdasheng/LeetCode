package main

import (
	"fmt"
	"time"
)

func main() {
	//strChan := make(chan string, 3)
	//strChan <- "a"
	//strChan <- "b"
	//strChan <- "c"
	//fmt.Println(strChan)

	strChan := make(chan string, 5)
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)

	go func() {
		<-syncChan1
		fmt.Println("go1 睡眠 1s，让下面的 go2 处于阻塞状态，只至 strChan 消费一个值后唤醒 go2")
		time.Sleep(time.Second)
		for {
			if elem, ok := <-strChan; ok {
				fmt.Println(elem)
			} else {
				break
			}
		}
		syncChan2 <- struct{}{}
	}()

	go func() {
		for _, elem := range []string{"a", "b", "c", "d"} {
			strChan <- elem
			if elem == "c" {
				syncChan1 <- struct{}{}
				fmt.Println("syncChan1 发送信号")
			}
		}
		fmt.Println("go2 睡眠 2s，防止接收方还没有全部接收完通道就关闭了")
		time.Sleep(2 * time.Second)
		close(strChan)
		fmt.Println("关闭了 strChan 通道")
		syncChan2 <- struct{}{}
	}()
	<-syncChan2
	<-syncChan2
	close(syncChan1)
	close(syncChan2)
}
