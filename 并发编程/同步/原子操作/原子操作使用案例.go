package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var num int32 = 32
	total := 10 // 总并发数
	var wg sync.WaitGroup
	wg.Add(total * 2)
	for i := 0; i < total; i++ {
		go func() {
			fmt.Printf("进行了 +3 操作，num=%d\n", num)
			func() {
				time.Sleep(time.Microsecond * 500)
				atomic.AddInt32(&num, 3)
			}()
			wg.Done()
		}()
	}
	for i := 0; i < total; i++ {
		go func() {
			fmt.Printf("进行了 -2 操作，num=%d\n", num)
			func() {
				time.Sleep(time.Microsecond * 500)
				atomic.AddInt32(&num, -2)
			}()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(num)
}

func add(v int32) int32 {
	time.Sleep(time.Microsecond * 500) // 模拟异步逻辑处理操作
	return v + 3
}

func sub(v int32) int32 {
	time.Sleep(time.Microsecond * 500) // 模拟异步逻辑处理操作
	return v - 2
}
