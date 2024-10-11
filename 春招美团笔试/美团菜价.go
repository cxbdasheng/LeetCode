package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	var sum int = 0
	sub := 0
	red := 0
	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		n, _ = strconv.Atoi(input.Text())
	}
	if input.Scan() {
		s := input.Text()
		arr := strings.Split(s, " ")
		for i := 0; i < n; i++ {
			now, _ := strconv.Atoi(arr[i])
			sum = sum + now
		}
	}

	if input.Scan() {
		s := input.Text()
		arr := strings.Split(s, " ")
		if len(arr) == 2 {
			sub, _ = strconv.Atoi(arr[0])
			red, _ = strconv.Atoi(arr[1])
		}
		if len(arr) == 1 {
			sub, _ = strconv.Atoi(arr[0])
		}

	}
	fmt.Println(math.MaxInt)
	fmt.Printf("sum：%d----sub：%d-----red:%d \n", sum, sub, red)
	// 总和大于满减+红包
	if sum >= sub+red {
		fmt.Println(sum - sub - red)
		return
	} else {
		// 总和大于红包则减红包
		if sum-red > 0 {
			fmt.Println(sum - red)
		} else {
			//红包大于0
			fmt.Println(0)
		}
	}
}
