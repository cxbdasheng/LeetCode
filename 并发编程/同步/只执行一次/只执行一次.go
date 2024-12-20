package main

import (
	"fmt"
	"sync"
)

func main() {
	var one sync.Once
	var count int
	for i := 0; i < 10; i++ {
		one.Do(func() {
			count++
		})
	}
	fmt.Println(count)
}
