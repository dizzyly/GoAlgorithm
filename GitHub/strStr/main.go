package main

import "fmt"

/*
	给定一个  haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从 0 开始)。
	如果不存在，则返回  -1。
*/

func strStr	(haystack string ,needle string) int {
	if needle == "" {
		return 0
	}
	var i, j int
	for i = 0; i <= len(haystack)-len(needle); i++ {
		for j = 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		// 判断字符串长度是否相等
		if len(needle) == j {
			return i
		}
	}
	return -1

}

func main() {
	fmt.Println(strStr("", ""))
}