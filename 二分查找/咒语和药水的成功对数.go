package main

import "sort"

func main() {

}
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr2)
	for i, num := range arr1 {
		sort.SearchInts(arr2, num+1)
	}
}
