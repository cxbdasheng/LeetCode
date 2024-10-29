package main

import (
	"fmt"
	"sync"
	"time"
)

// mutex.Lock() 多次操作会阻塞；mutex.Unlock() 逆操作会引起报错
func main() {
	var mutex sync.Mutex
	fmt.Println("Lock the lock. (main)")
	mutex.Lock()
	fmt.Println("The lock is locked. (main)")
	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Printf("Lock the lock. (g%d)\n", i)
			mutex.Lock()
			fmt.Printf("The lock is locked. (g%d)\n", i)
		}(i)
	}
	// 让相关协程可以执行
	time.Sleep(time.Second)
	fmt.Println("Unlock the lock. (main)")
	mutex.Unlock()
	fmt.Println("The lock is Unlock. (main)")
	// 让相关协程可以执行
	time.Sleep(time.Second)
}
