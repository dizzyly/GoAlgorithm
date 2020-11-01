package main

import (
	"GoAlgorithm/LeetCode/LinkedList"
	"fmt"
)

/*
	请判断一个链表是否为回文链表
*/
func main() {
	list := LinkedList.NewLinkedList()
	list.InsertToTail(1)
	list.InsertToTail(2)
	//list.InsertToTail(3)
	//list.InsertToTail(2)
	//list.InsertToTail(1)
	fmt.Println(isPalindrome(list.Head))
}

func isPalindrome(head *LinkedList.ListNode) bool {
	if head == nil {
		return false
	}
	var nums []interface{}
	for head != nil {
		nums = append(nums,head.Val)
		head = head.Next
	}
	//nums = nums[1:]
	fmt.Println(nums)
	for i := 0; i < len(nums)/2; i++ {
		if nums[i] != nums[len(nums)-i-1] {
			return false
		}

	}
	return true
}