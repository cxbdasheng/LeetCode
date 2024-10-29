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
			fmt.Printf("Try to lock for reading... [%d]\n", i)
			rwm.RLock()
			fmt.Printf("Locked for reading. [%d]\n", i)
			time.Sleep(time.Second * 2)
			fmt.Printf("Try to unlock for reading... [%d]\n", i)
			rwm.RUnlock()
			fmt.Printf("Unlocked for reading. [%d]\n", i)
		}(i)
	}
	time.Sleep(time.Microsecond * 100)
	fmt.Println("Try to lock for writing...")
	rwm.Lock()
	fmt.Println("Lock for writing...")
}
