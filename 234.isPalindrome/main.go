package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var nums = []int{}

	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	var low, high = 0, len(nums) - 1
	for low <= high {
		if nums[low] == nums[high] {
			low++
			high--
		} else {
			return false
		}
	}
	return true
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}
	res := isPalindrome(head)
	fmt.Println(res)
}
