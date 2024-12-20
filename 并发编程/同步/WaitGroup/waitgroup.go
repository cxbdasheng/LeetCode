package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("第一个协程")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("第二个协程")
	}()
	wg.Wait()

}
