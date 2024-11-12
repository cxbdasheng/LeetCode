package main

import "fmt"

// aerfadraer
// raefadr
// raeeeee
func main() {
	var s string
	fmt.Scan(&s)
	left := 0
	sum := 0
	mp := make([]int, 3)
	mp[0] = 0
	mp[1] = 0
	mp[2] = 0
	n := len(s)
	for i := 0; i < n; i++ {
		if s[i] == 'r' {
			mp[0]++
		} else if s[i] == 'e' {
			mp[1]++
		} else if s[i] == 'd' {
			mp[0] = 0
			mp[1] = 0
			continue
		}
		for mp[0] > 0 && mp[1] > 0 {
			j := i
			for ; j < n; j++ {
				if s[j] == 'd' {
					break
				}
			}
			sum += j - i
			if s[left] == 'r' {
				mp[0]--
			} else if s[left] == 'e' {
				mp[1]--
			}
			left++

		}
	}
	fmt.Println(sum)
}
