package main

import (
	"fmt"
	"sync"
	"time"
)

// 并发安全是指：程序能够正常运行，且可以得到正常运行的结果。
// 使用两个协程同时写会报错
//func main() {
//	s := make(map[int]int)
//	for i := 0; i < 100; i++ {
//		go func(i int) {
//			s[i] = i
//		}(i)
//	}
//	for i := 0; i < 100; i++ {
//		go func(i int) {
//			fmt.Printf("map 第 %d 个元素的值是 %d\n", i, s[i])
//		}(i)
//	}
//	time.Sleep(1 * time.Second)
//
//	// 上述会直接报错
//}

// 安全的 map
func main() {
	var mutex sync.RWMutex
	s := make(map[int]int)
	for i := 0; i < 100; i++ {
		go func(i int) {
			mutex.Lock()
			s[i] = i
			mutex.Unlock()
		}(i)
	}
	for i := 0; i < 100; i++ {
		go func(i int) {
			mutex.RLock()
			fmt.Printf("map 第 %d 个元素的值是 %d\n", i, s[i])
			mutex.RUnlock()
		}(i)
	}
	time.Sleep(1 * time.Second)
}
