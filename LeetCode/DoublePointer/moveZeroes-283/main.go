package main

import "fmt"

/*
	给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
*/

func moveZeroes(nums []int)  {
	var slow, fast int
	for fast = 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
	}
	for i := slow; i < len(nums); i++ {
		nums[i] = 0
	}

}

func main() {
	nums := []int{1,2,0,0,0,3,4,0,0,34,4,2,0,0}
	moveZeroes(nums)
	fmt.Println(nums)
}
