package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val:   4,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
		Right: &TreeNode{Val: 7, Left: nil, Right: nil},
	}
	fmt.Println(searchBST(root, 5))
}

func searchBST(root *TreeNode, val int) *TreeNode {
	var node = root
	for {
		if node == nil {
			return &TreeNode{}
		}
		if node.Val > val {
			node = node.Left
		} else if node.Val < val {
			node = node.Right
		} else if node.Val == val {
			return node
		}
	}
	return &TreeNode{}
}
