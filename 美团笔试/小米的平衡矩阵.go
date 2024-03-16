package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		n, _ = strconv.Atoi(input.Text())
	}
	t := make([][]int, n)
	for i := 0; i < n; i++ {
		t[i] = make([]int, n)
		if input.Scan() {
			str := input.Text()
			for j := 0; j < n; j++ {
				if len(str)-1 < j {
					t[i][j] = 0
					continue
				}
				t[i][j] = int(str[j]) - int('0')
			}
			fmt.Println(str)
		}
	}
	//strings.Split()
	//strings.Split()
}
