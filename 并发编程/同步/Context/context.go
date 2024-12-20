package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	coordinateWithContext()

	syncChan := make(chan struct{}, 1)
	go func() {
		for {
			select {
			case <-syncChan:
				fmt.Println("结束协程。。。")
				return
			default:
				fmt.Println("持续运行中")
			}
		}
	}()
	time.Sleep(time.Second)
	syncChan <- struct{}{}
	time.Sleep(time.Second * 2)

	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("结束协程。。。")
				return
			default:
				fmt.Println("持续运行中")
			}
		}
	}(ctx)
	time.Sleep(time.Second)
	cancel() // 发送通知
	time.Sleep(time.Second * 2)

}
func coordinateWithContext() {
	total := 30
	var num int64
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < total; i++ {
		go func(ctx context.Context) {
			atomic.AddInt64(&num, 1)
			<-ctx.Done()
		}(ctx)
	}
	time.Sleep(time.Second)
	cancel()
	fmt.Println(num)
}
