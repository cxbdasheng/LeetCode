package main

import (
	"errors"
	"fmt"
)

func main() {
	//index := 3
	//ay := [2]int{1, 2}
	//_ = ay[index]
	output()
	fmt.Println(11)
}

func output() {
	inner()
}

func inner() {
	panic(errors.New("预定义故障"))
}
