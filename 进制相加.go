package main

import "fmt"

func main() {
	//
	a := "11"
	b := "110"
	i, j := len(a)-1, len(b)-1
	var num = 0
	var s = ""
	for {
		if i >= 0 {
			num += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			num += int(b[j] - '0')
			j--
		}
		s = fmt.Sprintf("%d%s", num%2, s)
		num = num / 2
		if j < 0 && i < 0 {
			if num == 1 {
				s = "1" + s
			}
			break
		}
	}
	fmt.Println(s)
}
