//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//示例：
//
//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807

package main

import (
	"encoding/json"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var newListNode = &ListNode{}
	curr := newListNode
	if l1 != nil && l2 == nil {
		return l1
	}

	if l1 == nil && l2 != nil {
		return l2
	}
	var carry = 0
	for l1 != nil || l2 != nil {
		var (
			x = 0
			y = 0
		)
		if l1 != nil {
			x = l1.Val
		}

		if l2 != nil {
			y = l2.Val
		}
		sum := x + y + carry
		carry = sum / 10
		curr.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		curr = curr.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry > 0 {
		curr.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}

	return newListNode.Next
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	res := addTwoNumbers(l1, l2)
	bytes, _ := json.Marshal(res)
	fmt.Println(string(bytes))
}
