package main

import "GoAlgorithm/LeetCode/LinkedList"

/*
	将两个升序链表合并为一个新的 升序 链表并返回。
	新链表是通过拼接给定的两个链表的所有节点组成的。
*/


func main() {
	l1 := LinkedList.NewLinkedList()
	l2 := LinkedList.NewLinkedList()
	l1.InsertToTail(1)
	l1.InsertToTail(2)
	l1.InsertToTail(4)
	l2.InsertToTail(1)
	l2.InsertToTail(3)
	l2.InsertToTail(4)
	l3 := LinkedList.NewLinkedList()
	l3.Head = mergeTwoLists(l1.Head.Next, l2.Head.Next)
	l3.PrintList()
}
//方法一递归
//func mergeTwoLists(l1 *LinkedList.ListNode, l2 *LinkedList.ListNode) *LinkedList.ListNode {
//	if l1 == nil {
//		return l2
//	}
//	if l2 == nil {
//		return l1
//	}
//	if l1.Val < l2.Val {
//		l1.Next = mergeTwoLists(l1.Next,l2)
//		return l1
//	}
//	l2.Next = mergeTwoLists(l1, l2.Next)
//	return l2
//}
//方法二迭代
func mergeTwoLists(l1 *LinkedList.ListNode, l2 *LinkedList.ListNode) *LinkedList.ListNode {
	cur := &LinkedList.ListNode{}
	head := cur
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next

		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return head
}