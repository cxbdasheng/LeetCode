package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 输入数据有多组, 每行表示一组输入数据。
// 每行不定有n个整数，空格隔开。(1 <= n <= 100)。
// 1 2 3
// 4 5
// 0 0 0 0 0
func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		parts := strings.Split(sc.Text(), " ")
		var sum int
		for i := range parts {
			n, _ := strconv.Atoi(parts[i])
			sum += n
		}
		fmt.Println(sum)
	}
}
