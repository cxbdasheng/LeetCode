package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	a := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	val := tree(a)
	fmt.Println(Max(val[0], val[1]))
}

func tree(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	left := tree(root.Left)
	right := tree(root.Right)

	check := left[0] + right[0] + root.Val
	noCheck := Max(left[0], left[1]) + Max(right[1], right[0])

	return []int{noCheck, check}
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
