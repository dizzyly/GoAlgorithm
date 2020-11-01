package main

import (
	"GoAlgorithm/LeetCode/Tree"
	"fmt"
)

/*
	给定一个二叉树，它的每个结点都存放一个0-9的数字，每条从根到叶子节点的路径都代表一个数字。
	例如，从根到叶子节点路径 1->2->3 代表数字 123。
	计算从根到叶子节点生成的所有数字之和。
*/
func main() {
	root := Tree.CreateNode(1)
	root.Left = Tree.CreateNode(2)
	root.Right = Tree.CreateNode(3)
	root.Left.Left = Tree.CreateNode(4)
	root.Left.Left.Left = Tree.CreateNode(6)
	root.Left.Left.Right = Tree.CreateNode(7)
	root.Left.Right = Tree.CreateNode(5)
	root.Left.Right.Left = Tree.CreateNode(8)
	root.Left.Right.Right = Tree.CreateNode(9)
	fmt.Println(sumNumbers(root))
}

func sumNumbers(root *Tree.TreeNode) int {
	return dfs(root,0)
}
func dfs(root *Tree.TreeNode,prevSum int) int {
	if root == nil {
		return 0
	}
	sum := prevSum * 10 + root.Val
	if root.Left == nil && root.Right == nil {
		return  sum
	}
	return dfs(root.Left, sum) + dfs(root.Right, sum)
}