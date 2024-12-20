package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	count := 20  // 假设总共生产 20 个商品
	factory := 0 // 工厂中存在的商品
	cond := sync.NewCond(&mutex)

	for i := 0; i < count; i++ {
		go func() {
			mutex.Lock()
			factory = factory + 1
			fmt.Printf("生产了一个，工厂中剩余：%d\n", factory)
			mutex.Unlock()
			cond.Signal() // 发送唤醒通知
		}()
	}

	for i := 0; i < count; i++ {
		go func() {
			mutex.Lock()
			for factory <= 0 {
				cond.Wait() // 等待唤醒
			}
			factory = factory - 1
			fmt.Printf("消费了一个，工厂中剩余：%d\n", factory)
			mutex.Unlock()
		}()
	}
	time.Sleep(time.Second * 2)
	fmt.Println(factory)
}
