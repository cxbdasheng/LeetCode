package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	var nums []byte
	mp := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		for mp[s[i]] > 0 {
			nums = nums[1:]
			mp[s[i]]--
		}
		nums = append(nums, s[i])
		mp[s[i]]++
	}
	fmt.Println(string(nums))
}
