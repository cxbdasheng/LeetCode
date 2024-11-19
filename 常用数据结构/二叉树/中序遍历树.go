package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 8, Left: &TreeNode{Val: 8, Left: &TreeNode{Val: 9, Left: nil, Right: nil}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: nil, Right: nil}, Right: &TreeNode{Val: 7, Left: nil, Right: nil}}}, Right: &TreeNode{Val: 7, Left: nil, Right: nil}}
	fmt.Println(levelOrder(root))
}
func levelOrder(root *TreeNode) [][]int {
	var result [][]int

	if root == nil {
		return result
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		var vals []int
		l := len(queue)
		for i := 0; l < len(queue); i++ {
			head := queue[0]
			queue = queue[1:]
			vals = append(vals, head.Val)
			if head.Left != nil {
				queue = append(queue, head.Left)
			}
			if head.Right != nil {
				queue = append(queue, head.Right)
			}
		}
		result = append(result, vals)
	}
	return result
}
