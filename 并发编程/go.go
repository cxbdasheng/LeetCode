package main

import (
	"fmt"
	"time"
)

func main() {
	//go fmt.Println("hello world")
	//time.Sleep(time.Second)

	names := []string{"dasheng", "zhangsan", "lisi", "wangwu"}
	for _, name := range names {
		go func(s string) {
			fmt.Println(s)
		}(name)
	}
	time.Sleep(time.Second)
}
