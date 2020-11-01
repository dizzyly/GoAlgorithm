package main

import "fmt"

/*
	给你一个整数数组 arr，请你帮忙统计数组中每个数的出现次数。
	如果每个数的出现次数都是独一无二的，就返回 true；否则返回 false。
*/
func main() {
	arr := []int {1, 2, 3, 2, 1}
	fmt.Println(uniqueOccurrences(arr))
}

func uniqueOccurrences(arr []int) bool {
	var count = make(map[int]int)
	for _, v := range arr {
		count[v]++
	}
	times := map[int]struct{}{}
	for _, v := range count {
		times[v] = struct{}{}

	}
	//var nums []int
	//for _, v := range count {
	//	nums = append(nums, v)
	//}
	//for i := 0; i < len(nums) -1; i++ {
	//	for j:= i+1; j<len(nums); j++ {
	//		if nums[i] == nums[j] {
	//			return false
	//		}
	//	}
	//}
	return len(count) == len(times)
}