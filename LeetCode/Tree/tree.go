package Tree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func CreateNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
		Left: nil,
		Right: nil,
	}
}
