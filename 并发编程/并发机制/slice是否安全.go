package main

import (
	"fmt"
	"sync"
)

// 导致值没有这么多值
//
//	func main() {
//		var wg sync.WaitGroup
//		var s []int
//		for i := 0; i < 100; i++ {
//			wg.Add(1)
//			go func(i int) {
//				s = append(s, i)
//				wg.Done()
//			}(i)
//		}
//		wg.Wait()
//		fmt.Println(len(s), s)
//	}

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	var s []int
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			mutex.Lock()
			s = append(s, i)
			mutex.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Println(len(s), s)
}
