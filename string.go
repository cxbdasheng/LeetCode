package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%q\t%c \n", i, r)
		i += size
	}

	for i, r := range s {
		fmt.Printf("%q\t%c \n", i, r)
	}
	// 字符串转int
	// strconv.Atoi()

	c := "leetcode"
	fmt.Println(c[0:1])
}
