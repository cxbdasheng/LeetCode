package main

import (
	"container/list"
	"fmt"
	"math"
)

func main() {
	//var s = "110"
	fmt.Println(IntToBin(8))
	//IntToBin(6)

}

func StrToInt(s string) int {
	l := len(s)
	var num = 0
	for i := 0; i < l; i++ {
		if s[i] == '0' {
			num += int(0 * math.Pow(2, float64(l-i-1)))
		} else if s[i] == '1' {
			num += int(1 * math.Pow(2, float64(l-i-1)))
		}
	}
	return num
}

func IntToBin(num int) string {
	list.New()
	var binary []int
	var s string
	for num != 0 {
		binary = append(binary, num%2)
		num = num / 2
	}
	fmt.Println(binary)
	if len(binary) == 0 {
		s += "0"
	} else {
		for i := len(binary) - 1; i >= 0; i-- {
			s = fmt.Sprintf("%s%d", s, binary[i])
		}
	}
	return s
}
