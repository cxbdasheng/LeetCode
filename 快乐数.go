package main

import "fmt"

func main() {
	fmt.Println(isHappy(19))
}
func isHappy(n int) bool {
	m := make(map[int]bool)
	for {
		if n == 1 {
			break
		}
		if m[n] == true {
			return false
		}
		arr := getArr(n)
		m[n] = true
		n = 0
		for i := 0; i < len(arr); i++ {
			n += arr[i] * arr[i]
		}
	}
	return true
}

func getArr(n int) []int {
	var ar []int
	for n >= 10 {
		ar = append(ar, n%10)
		n = n / 10
	}
	ar = append(ar, n)
	return ar
}
