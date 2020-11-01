package main

import (
	"fmt"
)

/*
	给出两个非空的链表用来表示两个非负的整数。其中，它们各自的位数是按照逆序的方式存储的，并且它们的每个节点只能存储一位数字。
	如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
	您可以假设除了数字 0 之外，这两个数都不会以 0开头。
	示例：
	输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
	输出：7 -> 0 -> 8
	原因：342 + 465 = 807

*/

func main() {
	l1 := NewLinkedList()
	l2 := NewLinkedList()
	l1.InsertToTail(2)
	l1.InsertToTail(4)
	l1.InsertToTail(3)
	l2.InsertToTail(5)
	l2.InsertToTail(6)
	l2.InsertToTail(4)
	l3 := NewLinkedList()
	l3.Head.Next = addTwoNumbers(l1.Head.Next, l2.Head.Next)
	l3.PrintList()
}

func addTwoNumbers(l1, l2 *ListNode)  *ListNode {
	var tail *ListNode
	var head *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}

type ListNode struct {
	 Val int
	 Next *ListNode
}
type LinkedList struct {
	Head *ListNode
	Length uint
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}
func NewListNode(v int) *ListNode {
	return &ListNode{v, nil}
}

func (l *LinkedList) InsertAfter(cur *ListNode, val int) bool {
	if cur == nil {
		return false
	}
	newNode := NewListNode(val)
	newNode.Next = cur.Next
	cur.Next = newNode
	l.Length++
	return true
}

func (l *LinkedList) InsertToTail(v int) bool {
	cur := l.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	return l.InsertAfter(cur, v)

}
func (list *LinkedList) PrintList()  {
	cur := list.Head.Next
	format := ""
	for cur != nil {
		format += fmt.Sprintf("%+v",cur.GetValue())
		cur = cur.Next
		if cur != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}
func (node *ListNode) GetValue() interface{} {
	return node.Val
}


