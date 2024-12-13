package main

import "fmt"

func main() {
	ch := make(chan int, 5)
	syncChan := make(chan struct{}, 2)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
		syncChan <- struct{}{}
	}()

	go func() {
		for elem := range ch {
			fmt.Println(elem)
		}
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}
