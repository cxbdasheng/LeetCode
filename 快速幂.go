package main

import "fmt"

func main() {
	fmt.Println(fastPower(2, 100000000))
}
func fastPower(base, pow int) int {
	if pow == 0 {
		return 1
	}
	result := 1
	for pow != 0 {
		if pow&1 == 1 {
			result *= base
		}
		base *= base
		pow = pow >> 1
	}

	return result
}
