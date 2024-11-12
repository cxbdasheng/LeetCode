package main

import "fmt"

func main() {
	fmt.Println(maxPower("tourist"))
}

func maxPower(s string) int {
	m := 0
	sum := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			sum++
			if m < sum {
				m = sum
			}
		} else {
			sum = 1
		}
	}
	return m
}
