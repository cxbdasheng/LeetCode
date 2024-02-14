package main

import "fmt"

var m = make(map[string]int)

type address struct {
	hostname string
	port     int
}

func main() {
	//m := make(map[string]bool)
	//m["sss"] = true
	//fmt.Println(m)

	c := []string{"azs", "aaa"}
	fmt.Println(k(c))
	Add(c)
	fmt.Println(m)
	fmt.Println(Count(c))
	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
	fmt.Println(hits)
}
func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

func Add(list []string) {
	m[k(list)]++
}

func Count(list []string) int {
	return m[k(list)]
}
