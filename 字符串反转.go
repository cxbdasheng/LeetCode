package main

import "fmt"

func main() {
	var s = "abssaz"
	ss := []rune(s)
	i := 0
	j := len(s) - 1
	for i < j {
		ss[i], ss[j] = ss[j], ss[i]
		i++
		j--
	}
	fmt.Println(string(ss))
}
