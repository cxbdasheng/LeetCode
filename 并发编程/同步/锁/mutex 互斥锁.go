package main

import (
	"fmt"
	"sync"
	"time"
)

// mutex.Lock() 多次操作会阻塞；mutex.Unlock() 逆操作会引起报错
func main() {
	var mutex sync.Mutex
	fmt.Println("锁住这个锁. (主锁)")
	mutex.Lock()
	fmt.Println("这个锁已锁住. (主锁)")
	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Printf("锁住这个锁. (协程%d)\n", i)
			mutex.Lock()
			fmt.Printf("这个锁已锁住. (协程%d)\n", i)
		}(i)
	}
	// 让相关协程可以执行
	time.Sleep(time.Second)
	fmt.Println("解锁这个锁. (主锁)")
	mutex.Unlock()
	fmt.Println("已解锁这个锁. (主锁)")
	// 让相关协程可以执行
	time.Sleep(time.Second)
}
