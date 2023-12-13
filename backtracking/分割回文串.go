package main

import "fmt"

var result [][]string
var path []string

func main() {
	//fmt.Println(partition("aab"))
	fmt.Println(partition("aab"))
}

func partition(s string) [][]string {
	backtracking(s, 0)
	return result
}

func backtracking(s string, start int) {
	if start == len(s) {
		tmp := make([]string, len(path))
		copy(tmp, path)
		result = append(result, tmp)
		return
	}
	for i := start; i < len(s); i++ {
		if IsChild(s[start : i+1]) {
			path = append(path, s[start:i+1])
			backtracking(s, i+1)
			path = path[:len(path)-1]
		}
	}
}

func IsChild(s string) bool {
	if s == "" {
		return false
	}
	l := len(s)

	for i, j := 0, l-1; i < j; i++ {
		if s[i] == s[j] {
			j--
		} else {
			return false
		}
	}
	return true
}
