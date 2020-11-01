package main

/*
	给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
*/
type ListNode struct {
	Val interface{}
	Next *ListNode
}
type LinkList struct {
	head *ListNode
	length uint
}
func NewListNode (v interface{}) *ListNode {
	return &ListNode{v,nil}
}
func NewList() *LinkList {
	return &LinkList{
		head: NewListNode(0),
		length: 0,
	}
}
func (list *LinkList)InsertToTail (v interface{}) bool {
	cur := list.head
	for cur.Next != nil {
		cur = cur.Next
	}
	return list.InsertAfter(cur, v)
}
func (list LinkList) InsertAfter(p *ListNode,v interface{}) bool {
	if p == nil || list.head == p {
		return false
	}
	cur := list.head.Next
	pre := list.head
	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.Next
	}
	if cur == nil {
		return false
	}
	newNode := NewListNode(v)
	newNode.Next = cur
	pre.Next = newNode
	list.length++
	return true
}

func deleteDuplicates(head *ListNode) *ListNode {
	 if head == nil {
	 	return head
	 }
	 slow := head
	 fast := head
	 for fast != nil {
	 	 if fast.Val != slow.Val {
	 	 	slow.Next = fast
	 	 	slow = slow.Next
		 }
		 fast = fast.Next
	 }
	 slow.Next = nil
	 return head
}

func main() {

}
