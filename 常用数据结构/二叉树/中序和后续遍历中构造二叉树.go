package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(buildTree([]int{9, 15, 7, 20, 3}, []int{9, 3, 15, 20, 7}))
}

func buildTree(postorder []int, inorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootVal}
	if len(postorder) == 1 {
		return root
	}
	var index int
	for index = 0; index < len(inorder); index++ {
		if inorder[index] == rootVal {
			break
		}
	}
	// 切中序数组
	leftInorder := inorder[:index]
	rightInorder := inorder[index+1:]
	// postorder 舍弃末尾元素
	postorder = postorder[:len(postorder)-1]
	// 切割后序数组
	leftPostorder := postorder[:len(leftInorder)]
	rightPostorder := postorder[len(leftInorder):]
	root.Left = buildTree(leftInorder, leftPostorder)
	root.Right = buildTree(rightInorder, rightPostorder)

	return root
}
