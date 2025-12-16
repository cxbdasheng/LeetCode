package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(answerQueries([]int{736411, 184882, 914641, 37925, 214915}, []int{331244, 273144, 118983, 118252, 305688, 718089, 665450}))
}
func answerQueries(nums []int, queries []int) []int {
	sort.Ints(nums)
	sumNums := make([]int, 4, len(nums))
	sumNums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sumNums[i] = nums[i] + sumNums[i-1]
	}
	for i := 0; i < len(queries); i++ {
		index := sort.SearchInts(sumNums, queries[i]+1)
		queries[i] = index
	}
	return queries

}
