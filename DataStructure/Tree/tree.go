package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func main() {
	root := CreateNode(1)

	root.left = CreateNode(2)
	root.right = CreateNode(3)
	root.left.left = CreateNode(4)
	root.left.right = CreateNode(5)
	root.right.left = CreateNode(6)
	//root.PreOrder(root)
	//root.InOrder(root)
	root.PostOrder(root)
	//root.GetLeafNode(root)
	//root.preorderTraversal(root)
	//root.inorderTraversal(root)
	root.postorderTraversal(root)
}

func CreateNode(val int) *Node {
	return &Node{
		value: val,
		left:  nil,
		right: nil,
	}
}

//前序遍历
func (node *Node) PreOrder(n *Node) {
	if n != nil {
		fmt.Printf("%d ", n.value)
		node.PreOrder(n.left)
		node.PreOrder(n.right)
	}
}

//中序遍历
func (node *Node) InOrder(n *Node) {
	if n != nil {
		node.InOrder(n.left)
		fmt.Printf("%d ", n.value)
		node.InOrder(n.right)
	}
}

//后序遍历
func (node *Node) PostOrder(n *Node) {
	if n != nil {
		node.PostOrder(n.left)
		node.PostOrder(n.right)
		fmt.Printf("%d ", n.value)
	}
}

//打印叶子节点
func (node *Node) GetLeafNode(n *Node) {
	if n != nil {
		if n.left == nil && n.right == nil {
			fmt.Printf("%d ", n.value)
		}
		node.GetLeafNode(n.left)
		node.GetLeafNode(n.right)
	}
}

//前序非递归遍历
func (node *Node) preorderTraversal(root *Node) {
	// 非递归
	if root == nil {
		fmt.Println("空树")
	}
	result := make([]int, 0)
	stack := make([]*Node, 0)

	for root != nil || len(stack) != 0 {
		for root != nil {
			// 前序遍历，所以先保存结果
			result = append(result, root.value)
			stack = append(stack, root)
			root = root.left
		}
		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.right
	}
	fmt.Println(result)
}

//非递归中序遍历
func (node *Node) inorderTraversal(root *Node) {
	if root == nil {
		fmt.Println("空树")
	}
	result := make([]int, 0)
	stack := make([]*Node, 0)

	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.value)
		root = node.right
	}
	fmt.Println(result)

}

//非递归后序遍历
func (node Node) postorderTraversal(root *Node) {
	if root == nil {
		fmt.Println("空树")
	}
	var lastVisit *Node
	result := make([]int, 0)
	stack := make([]*Node, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.left
		}
		// 这里先看看，先不弹出
		node := stack[len(stack)-1]
		// 根节点必须在右节点弹出之后，再弹出
		if node.right == nil || node.right == lastVisit {
			stack = stack[:len(stack)-1] // pop
			result = append(result, node.value)
			// 标记当前这个节点已经弹出过
			lastVisit = node
		} else {
			root = node.right
		}
	}
	fmt.Println(result)

}
