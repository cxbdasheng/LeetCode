package main

import "fmt"

func main() {
	var num uint32
	var count int
	num = 00000000000000000000000000001011
	for i := 0; i <= 32; i++ {
		if (num & (1 << i)) != 0 {
			count++
		}
	}
	fmt.Println(count)
}
