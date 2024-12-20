package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	t0 := time.Now()
	go func() {
		time.Sleep(1000 * time.Millisecond)
		cancel()
	}()
	<-ctx.Done()
	t1 := time.Now()
	fmt.Println(t1.Sub(t0).Milliseconds(), ctx.Err())
}
