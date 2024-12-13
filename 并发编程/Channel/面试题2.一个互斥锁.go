package main

import (
	"fmt"
	"time"
)

func main() {
	lock := make(chan struct{}, 1)
	lock <- struct{}{}

	go func() {
		<-lock
		fmt.Println("访问共享资源")
		lock <- struct{}{}
	}()

	go func() {
		<-lock
		fmt.Println("访问共享资源")
		lock <- struct{}{}
	}()
	time.Sleep(time.Second * 1)
}
