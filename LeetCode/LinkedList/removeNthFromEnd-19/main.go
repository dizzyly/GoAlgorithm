package main

import "fmt"

/*
	给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
*/

func main() {
	l1 := NewLinkedList()
	l1.InsertToTail(1)
	l1.InsertToTail(2)
	l1.InsertToTail(3)
	l1.InsertToTail(4)
	l1.InsertToTail(5)
	l1.PrintList()
	l2 := NewLinkedList()
	l2.Head.Next = removeNthFromEnd(l1.Head.Next, 5)
	l2.PrintList()
}
//方法二

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := NewListNode(0)
	dummy.Next = head
	fast, slow := head, dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next

}
//方法一
//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	var length int
//	cur := head
//	for ; cur != nil; cur = cur.Next {
//		length++
//
//	}
//	cur = head
//	if length == n {
//		return cur.Next
//	}
//	now :=  length - n
//
//	for i := 1; i < now; i++ {
//		cur = cur.Next
//	}
//	if cur.Next.Next == nil {
//		cur.Next = nil
//	} else {
//		cur.Next = cur.Next.Next
//	}
//	return head
//}

type ListNode struct {
	Val int
	Next *ListNode
}
type LinkedList struct {
	Head *ListNode
	Length int
}

func NewListNode(val int) *ListNode {
	return  &ListNode{val, nil}
}
func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0),0}
}

func (l *LinkedList) InsertToTail(v int) bool {
	cur := l.Head

	for cur.Next != nil {
		cur = cur.Next
	}
	newNode := NewListNode(v)
	cur.Next = newNode
	l.Length++
	return true
}
func (l *LinkedList) PrintList()  {
	cur := l.Head.Next
	format := ""
	for cur != nil {
		format += fmt.Sprintf("%+v",cur.Val)
		cur = cur.Next
		if cur != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}
