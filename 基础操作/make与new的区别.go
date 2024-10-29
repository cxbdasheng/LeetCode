package main

import "fmt"

func main() {
	a := new([]int)
	b := make([]int, 5)
	*a = append(*a, 1)
	fmt.Println(a, b)
}
