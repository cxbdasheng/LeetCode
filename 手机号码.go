package main

import (
	"fmt"
	"github.com/xluohome/phonedata"
)

func main() {
	pr, err := phonedata.Find("18957509123")
	if err != nil {
		panic(err)
	}
	fmt.Print(pr)
}
