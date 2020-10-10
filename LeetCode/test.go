package main

import "fmt"

func main() {
	s := []string{
		"hello",
		"helloorde",
		"dheld",
	}
	pre := s[0]
	for i := 1; i < len(s); i++ {
		 pre = lcp(pre,s[i])
		 if len(pre) == 0 {
		 	break
		 }
	}
	fmt.Println(pre)
}
func lcp(pre, s string) string {
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
	}else {
		return x
	}
}