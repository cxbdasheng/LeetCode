package main

import (
	"fmt"
	"sync"
)

func main() {
	pool := sync.Pool{New: func() interface{} {
		return make([]int, 0)
	}}
	fmt.Println(pool.Get().([]int))
	fmt.Println(pool.Get().([]int))
	fmt.Println(pool.Get().([]int))
	fmt.Println(pool.Get().([]int))
}
