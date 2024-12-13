package main

import "fmt"

func main() {
	strChan := make(chan string, 5)
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)

	go func() {
		<-syncChan1
		for {
			if elem, ok := <-strChan; ok {
				fmt.Println(elem)
			} else {
				break
			}
		}
		syncChan2 <- struct{}{}
	}()
	go func() {
		for _, elem := range []string{"a", "b", "c", "d"} {
			strChan <- elem
		}
		close(strChan)
		syncChan1 <- struct{}{}
		syncChan2 <- struct{}{}
	}()
	<-syncChan2
	<-syncChan2
}
