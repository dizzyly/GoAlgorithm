package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {

	var min int
	for root != nil {
		Leftnode := root.Left
		Rightnode := root.Right
		var leftDistence, rightDistence int
		var lfFlag, rtflag bool
		if Leftnode != nil {
			lfFlag = true
			leftDistence = root.Val - Leftnode.Val
		}
		if Rightnode != nil {
			rtflag = true
			rightDistence = Rightnode.Val - root.Val
		}
		if lfFlag && rtflag && leftDistence < rightDistence{
			min = leftDistence
		}
		getMinimumDifference(root.Left)
		getMinimumDifference(root.Right)
	}
	return min
}
func creatNode(val int) *TreeNode  {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}
func main() {
	root := creatNode(1)
	root.Right = creatNode(3)
	root.Right.Left = creatNode(2)
	fmt.Println(getMinimumDifference(root))
}
