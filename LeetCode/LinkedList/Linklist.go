package LinkedList

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}
type LinkedList struct {
	Head *ListNode
	Length uint
}
func NewListNode (v int) *ListNode {
	return &ListNode{v,nil}
}
func NewLinkedList () *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}
func (node *ListNode) GetNext() *ListNode {
	return node.Next
}
func (node *ListNode) GetValue() int {
	return node.Val
}
//在某个节点后面插入节点
func (list *LinkedList) InsertAfter(p *ListNode, v int) bool {
	if p == nil {
		return false
	}
	newNode := NewListNode(v)
	oldNode := p.Next
	p.Next = newNode
	newNode.Next = oldNode
	list.Length++
	return true
}
//在某个节点前面插入节点
func (list *LinkedList) InsertBefore(p *ListNode, v int) bool {
	if p == nil || list.Head == p {
		return false
	}
	cur := list.Head.Next
	pre := list.Head
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
	list.Length++
	return true
}
//在链表头部插入节点
func (list *LinkedList) InsertToHead( v int) bool {
	return list.InsertAfter(list.Head, v)
}
//在链表尾部插入节点
func (list *LinkedList) InsertToTail(v int) bool {
	cur := list.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	//newNode := NewListNode(v)
	//cur.Next = newNode
	//list.length++
	return list.InsertAfter(cur, v)
}
//通过索引查找节点
func (list *LinkedList) FindByIndex(index uint) *ListNode {
	if list.Length < index {
		return nil
	}
	cur := list.Head.Next
	var i uint = 0
	for ; i < index; i++ {
		 cur = cur.Next
	}
	return cur
}
//删除传入的节点
func (list *LinkedList) DeleteNode(p *ListNode) bool {
	if p == nil {
		return false
	}
	cur := list.Head.Next
	pre := list.Head
	for ; cur != p; {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.Next
	}
	if cur == nil {
		return false
	}
	pre.Next = cur.Next
	p = nil
	list.Length--
	return true
}
//打印链表
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
//单链表反转
//时间复杂度o(n)
func (list *LinkedList) ListReverse()  {
	if list.Head == nil || list.Head.Next == nil || list.Head.Next.Next == nil {
		return
	}
	//var pre *ListNode = nil
	//cur := list.head.Next
	//for cur != nil {
	//	temp := cur.Next
	//	cur.Next = pre
	//	pre = cur
	//	cur = temp
	//}
	//list.head.Next = pre

	cur := list.Head.Next.Next
	pre := list.Head.Next
	//var i uint = 0
	//for i = 1; i < list.length; i++ {
	//
	//	list.head.Next = cur
	//	pre.Next = cur.Next
	//	cur.Next = pre
	//	//cur = pre.Next
	//
	//}
	list.Head.Next = cur
	pre.Next = cur.Next
	cur.Next = pre
}
//func main() {
//	list := NewLinkedList()
//	list.InsertToTail(2)
//	list.InsertToHead(1)
//	list.InsertToTail(3)
//	list.InsertToTail(4)
//	list.PrintList()
//	list.ListReverse()
//	list.PrintList()
//
//}
