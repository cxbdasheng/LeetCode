package main

import "fmt"

// paectc
func main() {
	var s string
	fmt.Scan(&s)
	var sCopy []byte
	sCopy = []byte(s)
	var temp byte
	for i := 0; i < len(sCopy); i++ {
		temp = sCopy[i]
		sCopy = append(sCopy[:i], sCopy[i+1:]...)
		sCopy = append(sCopy, temp)
	}
	fmt.Println(string(sCopy))
}
