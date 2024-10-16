package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}
func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}
	if len(preorder) == 1 {
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
	// preorder 舍弃末尾元素
	preorder = preorder[1:]
	// 切割后序数组
	leftPostorder := preorder[:len(leftInorder)]
	rightPostorder := preorder[len(leftInorder):]
	root.Left = buildTree(leftPostorder, leftInorder)
	root.Right = buildTree(rightPostorder, rightInorder)

	return root
}
