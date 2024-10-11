package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	var n, q int
	if input.Scan() {
		s := input.Text()
		arr := strings.Split(s, " ")
		n, _ = strconv.Atoi(arr[0])
		q, _ = strconv.Atoi(arr[1])
	}
	path := make([]uint64, n)
	if input.Scan() {
		s := input.Text()
		arr := strings.Split(s, " ")
		for i := 0; i < n; i++ {
			now, _ := strconv.ParseUint(arr[i], 10, 64)
			path[i] = now
		}
	}
	for i := 0; i < q; i++ {
		if input.Scan() {
			s := input.Text()
			x, _ := strconv.Atoi(s)
			path = add(path, x)
		}
	}
	var sum uint64
	for i := 0; i < len(path); i++ {
		sum += path[i]
	}
	fmt.Println(sum)
}

func add(path []uint64, x int) []uint64 {
	for i := 0; i < len(path); i++ {
		if i == x-1 {
			continue
		}
		path[i] = path[i] * 2
	}
	return path
}
