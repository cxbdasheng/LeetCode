package main

import "fmt"

func main() {
	fmt.Println(maxVowels("abciiidef", 3))
}
func maxVowels(s string, k int) int {
	max := 0
	vowel := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			vowel++
			if vowel > max {
				max = vowel
			}
		}
		if i+1 < k {
			continue
		}
		if s[i+1-k] == 'a' || s[i+1-k] == 'e' || s[i+1-k] == 'i' || s[i+1-k] == 'o' || s[i+1-k] == 'u' {
			vowel--
		}
	}
	return max
}
