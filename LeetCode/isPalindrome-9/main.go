package main

import (
	"fmt"
	"strconv"
)

/*
	判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
*/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s :=strconv.FormatInt(int64(x), 10)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(-123211))
}
