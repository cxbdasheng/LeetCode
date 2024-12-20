package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Millisecond*1000))
	t1 := time.Now()
	defer cancel()
	// 睡眠 500
	time.Sleep(time.Millisecond * 100)
	// 子 context
	ctx1, cancel2 := context.WithDeadline(ctx, time.Now().Add(time.Millisecond*1200))
	t2 := time.Now()
	defer cancel2()
	// 阻塞等待
	<-ctx1.Done()
	t3 := time.Now()
	fmt.Println(t2.Sub(t1).Milliseconds(), t3.Sub(t2).Milliseconds(), ctx1.Err())
}
