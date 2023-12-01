package main

import "fmt"

func main() {
	ransomNote := "a"
	magazine := "baa"
	maps := make(map[string]int, 26)
	for i := 0; i < len(magazine); i++ {
		key := magazine[i]
		maps[string(key)]++
		//fmt.Printf("%T", key)
	}

	fmt.Println(maps)
	fmt.Println(ransomNote)
}
