package main

import (
	"fmt"
	"time"
)

func main() {
	syncChan := make(chan struct{}, 2)
	mpCh := make(chan map[string]int, 1)
	go func() {
		for {
			if elem, ok := <-mpCh; ok {
				elem["count"]++
			} else {
				break
			}
		}
		syncChan <- struct{}{}
	}()
	go func() {
		countMap := make(map[string]int)
		for i := 0; i < 5; i++ {
			mpCh <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("这个值的值是: %v \n", countMap)
		}
		close(mpCh)
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}
