package main

import "fmt"

/*
	给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
	有效字符串需满足：1.左括号必须用相同类型的右括号闭合。2.左括号必须以正确的顺序闭合。
	注意空字符串可被认为是有效字符串。

*/
func isValid(s string) bool {
	var stack []byte
	var m = map[byte]byte {
		'{':'}',
		'[':']',
		'(':')',
	}
	if s == "" {
		return true
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '{' || s[i] == '[' || s[i] == '('{

			stack = append(stack, m[s[i]])
		} else if len(stack) == 0{
			return  false

		} else {
			if s[i] == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			}else {
				return false
			}
		}
	}

		return true

}
func main() {
	fmt.Println(isValid("]"))
}
