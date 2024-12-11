package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	close(ch)
	elem, ok := <-ch
	fmt.Println(elem, ok)
	ch <- 2
	elem1, ok1 := <-ch
	fmt.Println(elem1, ok1)
}
