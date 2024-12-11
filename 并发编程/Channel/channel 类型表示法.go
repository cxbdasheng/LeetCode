package main

import "fmt"

func main() {
	//var intChan chan int
	type intChan chan int
	var a intChan
	fmt.Println(a)

	var acceptIntChan chan<- int
	fmt.Println(acceptIntChan)

	var sendIntChan <-chan int
	fmt.Println(sendIntChan)

	ch := make(chan int)
	fmt.Println(ch)

	ch2 := make(chan int)
	fmt.Println(ch2)
}
