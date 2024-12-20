package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写互斥，读完之后才能写，写完之后才能读。
func main() {
	var rwm sync.RWMutex
	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Printf("尝试锁定读 [%d]\n", i)
			rwm.RLock()
			fmt.Printf("读已锁定. [%d]\n", i)
			time.Sleep(time.Second * 2)
			fmt.Printf("尝试解锁读... [%d]\n", i)
			rwm.RUnlock()
			fmt.Printf("读已解锁定 [%d]\n", i)
		}(i)
	}
	time.Sleep(time.Microsecond * 100)
	fmt.Println("尝试锁定写")
	rwm.Lock()
	fmt.Println("写已锁定")
}
