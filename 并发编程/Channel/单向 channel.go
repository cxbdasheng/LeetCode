package main

import "fmt"

func main() {
	//ch := make(chan<- int, 1)
	//fmt.Println(ch)
	strChan := make(chan string, 5)
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)
	go receive(strChan, syncChan1, syncChan2)
	go send(strChan, syncChan1, syncChan2)
	<-syncChan2
	<-syncChan2
}

func receive(strChan <-chan string, syncChan1 <-chan struct{}, syncChan2 chan<- struct{}) {
	<-syncChan1
	for {
		if elem, ok := <-strChan; ok {
			fmt.Println(elem)
		} else {
			break
		}
	}
	syncChan2 <- struct{}{}
}

func send(strChan chan<- string, syncChan1 chan<- struct{}, syncChan2 chan<- struct{}) {
	for _, elem := range []string{"a", "b", "c", "d"} {
		strChan <- elem
		if elem == "c" {
			syncChan1 <- struct{}{}
			fmt.Println("syncChan1 发送信号")
		}
	}
	close(strChan)
	syncChan2 <- struct{}{}
}
