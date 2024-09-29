package main

import "fmt"

func main() {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	fmt.Println(single)
}
