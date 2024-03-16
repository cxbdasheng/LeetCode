package main

import "fmt"

func main() {
	var n, s, num int
	for {
		_, err := fmt.Scanf("%d", &n)
		if err != nil {
			return
		}
		if n == 0 {
			return
		}
		s = 0
		num = 0
		for i := 0; i < n; i++ {
			_, err := fmt.Scanf("%d", &num)
			if err != nil {
				return
			}
			s += num
		}
		fmt.Println(s)
	}
}
