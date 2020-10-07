package main

import "fmt"

type Node struct {
	value  int
	left  *Node
	right *Node
}

func main()  {
	root := CreateNode(1)

	root.left = CreateNode(2)
	root.right = CreateNode(3)
	root.left.left = CreateNode(4)
	root.left.right = CreateNode(5)
	root.right.left = CreateNode(6)
	//root.PreOrder(root)
	//root.InOrder(root)
	//root.PostOrder(root)
	root.GetLeafNode(root)
}

func CreateNode(val int) *Node {
	return &Node{
		value:  val,
		left:  nil,
		right: nil,
	}
}
//前序遍历
func (node *Node)PreOrder(n *Node) {
	if n != nil {
		fmt.Printf("%d ", n.value)
		node.PreOrder(n.left)
		node.PreOrder(n.right)
	}
}
//中序遍历
func (node *Node) InOrder(n *Node)  {
	if n != nil {
		node.InOrder(n.left)
		fmt.Printf("%d ", n.value)
		node.InOrder(n.right)
	}
}
//后序遍历
func (node *Node) PostOrder(n *Node)  {
	if n != nil {
		node.PostOrder(n.left)
		node.PostOrder(n.right)
		fmt.Printf("%d ", n.value)
	}
}
//打印叶子节点
func (node *Node) GetLeafNode(n *Node)  {
	if n != nil {
		if n.left == nil && n.right == nil {
			fmt.Printf("%d ", n.value)
		}
		node.GetLeafNode(n.left)
		node.GetLeafNode(n.right)
	}
}