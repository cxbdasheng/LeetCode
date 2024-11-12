package main

import "fmt"

// 1
// abcdefg
func main() {
	var t int
	fmt.Scan(&t)
	var s string
	for i := 0; i < t; i++ {
		fmt.Scan(&s)
		var temp string
		for i := 0; i < len(s); i++ {
			if SumBin(i+1)&1 == 1 {
				temp += string('A' + (s[i] - 'a'))
			} else {
				temp += string(s[i])
			}
		}
		fmt.Println(temp)
	}
}
func SumBin(n int) int {
	sum := 0
	for n != 0 {
		if n&1 == 1 {
			sum++
		}
		n >>= 1
	}
	return sum
}
