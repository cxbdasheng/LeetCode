package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var m = -1491
	var s = "20:50"
	var n = 0
	result := strings.Split(s, ":")
	minute, err := strconv.Atoi(result[1])
	if err != nil {
		fmt.Println("转换失败:", err)
		return
	}
	m = m + minute
	n = m / 60
	m = m % 60
	if m < 0 {
		m = 60 + m
		n = n - 1
	}
	hour, err := strconv.Atoi(result[0])
	if err != nil {
		fmt.Println("转换失败:", err)
		return
	}
	hour = hour + n
	hour = hour % 24
	if hour < 0 {
		hour = hour + 24
	}
	var min, h string
	if m <= 9 {
		min = fmt.Sprintf("0%d", m)
	} else {
		min = fmt.Sprintf("%d", m)
	}
	if hour <= 9 {
		h = fmt.Sprintf("0%d", hour)
	} else {
		h = fmt.Sprintf("%d", hour)
	}
	date := fmt.Sprintf("%s:%s", h, min)
	fmt.Println(date)
}
