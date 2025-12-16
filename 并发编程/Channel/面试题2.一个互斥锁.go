package main

import (
	"fmt"
)

func main() {
	lock := make(chan struct{}, 1)
	lock2 := make(chan struct{}, 2)
	lock <- struct{}{}

	go func() {
		<-lock
		fmt.Println("访问共享资源1")
		lock <- struct{}{}
		lock2 <- struct{}{}
	}()

	go func() {
		<-lock
		fmt.Println("访问共享资源2")
		lock <- struct{}{}
		lock2 <- struct{}{}
	}()
	<-lock2
	<-lock2
}
