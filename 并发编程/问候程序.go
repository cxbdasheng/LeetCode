package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	input, er := inputReader.ReadString('\n')
	if er != nil {
		panic(er)
	}
	fmt.Println(input)
}
