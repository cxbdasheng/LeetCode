package main

import "fmt"

// 2
// 1 1
// 4 2
func main() {
	var g uint64
	fmt.Scan(&g)
	var i uint64
	for i = 0; i < g; i++ {
		var n, k uint64
		fmt.Scan(&n, &k)
		if n == k {
			fmt.Println(1)
			continue
		}
		if n>>1 >= k {
			fmt.Println(k)
		} else {
			if n&1 == 1 {
				fmt.Println(n>>1 - (k - n>>1) + 1)
			} else {
				fmt.Println(n>>1 - (k - n>>1))
			}
		}
	}
}
