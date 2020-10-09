package main

import "fmt"

/*
	给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
	不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

*/

func main() {
	nums := []int{0,1,1,1,2,2,2,2,2,3,3,4,4,4,4,4}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	//for i := 0; i < len(nums) - 1; i++ {
	//	if nums[i] == nums[i+1] {
	//		nums = append(nums[:i],nums[i+1:]...)
	//		i--
	//	}
	//}
	//return len(nums)
	location := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		nums[location] = nums[i]
		location++
	}
	return location
}