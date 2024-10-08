package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

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
