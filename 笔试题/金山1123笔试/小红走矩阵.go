package main

import (
	"fmt"
	"strings"
)

// 3 3
// red
// der
// rdr
func main() {
	var n, m int
	fmt.Scan(&n, &m)
	nums := make([][]string, n)
	var str string
	for i := 0; i < n; i++ {
		fmt.Scan(&str)
		strArr := strings.Split(str, "")
		nums[i] = strArr
	}

}
