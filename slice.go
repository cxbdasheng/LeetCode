package main

import "fmt"

func main() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(cap(s))
	fmt.Println(remove(s, 2))
	fmt.Println(s)
	c := []string{"azs", "aaa"}
	fmt.Printf("%q", c)
	cq := fmt.Sprintf("%q", c)
	fmt.Println(cq)
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
