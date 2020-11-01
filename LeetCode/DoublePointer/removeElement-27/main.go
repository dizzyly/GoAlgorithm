package main

import "fmt"

/*
	给你一个数组 nums和一个值 val，你需要原地移除所有数值等于val的元素，并返回移除后数组的新长度。
	不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
	元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
	示例 1:
	给定 nums = [3,2,2,3], val = 3,
	函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。
	你不需要考虑数组中超出新长度后面的元素。

*/

func removeElement(nums []int, val int) int {
	var slow, fast int
	for fast = 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
func main() {
	nums := []int{1,1,2,2,5,2,6,2,6}
	fmt.Println(removeElement(nums,2))
}
