package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	total := 5
	for i := 0; i < total; i++ {
		go func(ctx context.Context, i int) {
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("结束协 %d。。。\n", i)
					return
				default:
					fmt.Printf("协程 %d，持续运行中\n", i)
					time.Sleep(1 * time.Second)
				}
			}
		}(ctx, i)
	}

	time.Sleep(2 * time.Second)
	cancel() // 发送通知
	time.Sleep(2 * time.Second)
}
