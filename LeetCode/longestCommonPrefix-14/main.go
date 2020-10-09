package main

import "fmt"
/*

	编写一个函数来查找字符串数组中的最长公共前缀。
	如果不存在公共前缀，返回空字符串 ""。

*/


func main() {
	strs := []string{
		"hell",
		"heel",
		"helel",

	}
	fmt.Println(longestCommonPrefix(strs))
}
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return " "
	}
	pre := strs[0]
	for i := 1; i < len(strs); i++ {
		pre = lcp(pre,strs[i])
		if len(pre) == 0 {
			return " "
		}
	}
	return pre

}
func lcp(pre, s string) string  {
	lenth := min(len(pre), len(s))
	var i int
	for i = 0; i < lenth; i++ {
		if pre[i] != s[i] {
			break
		}
	}
	return pre[:i]
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}