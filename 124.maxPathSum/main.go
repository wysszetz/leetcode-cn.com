package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

var Answer = INT_MIN

func maxPathSum(root *TreeNode) int {
	maxGain(root)
	return Answer
}

func maxGain(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := max(0, maxGain(root.Left))
	right := max(0, maxGain(root.Right))
	Answer = max(Answer, left+right+root.Val)
	return max(left, right) + root.Val
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	//root := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val:   2,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val:   3,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//}

	root := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	res := maxPathSum(root)
	fmt.Println(res)
}
