package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var beforeNode = head
	var nowNode = head.Next

	for nowNode != nil {
		if beforeNode.Val == nowNode.Val {
			beforeNode.Next = nowNode.Next
		} else {
			beforeNode = nowNode
		}
		nowNode = nowNode.Next
	}

	return head
}
func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
	}
	res := deleteDuplicates(head)
	fmt.Println(res)
}
