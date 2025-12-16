package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println(addTwoNumbers(&ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}, &ListNode{Val: 9}))
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	pre := dummy
	var isCarry bool
	for l1 != nil || l2 != nil {
		val := 0
		if isCarry {
			val++
			isCarry = false
		}
		if l1 != nil && l2 != nil {
			val = val + l1.Val + l2.Val
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 == nil {
			val = val + l2.Val
			l2 = l2.Next
		} else if l2 == nil {
			val = val + l1.Val
			l1 = l1.Next
		}
		if val >= 10 {
			isCarry = true
		}
		val = val % 10
		node := &ListNode{Val: val}
		pre.Next = node
		pre = node
	}
	if isCarry {
		node := &ListNode{Val: 1}
		pre.Next = node
	}
	return dummy.Next
}
