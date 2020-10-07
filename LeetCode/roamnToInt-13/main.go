package main

import (
	"fmt"
)
/*
	罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。
	字符          数值
	I             1
	V             5
	X             10
	L             50
	C             100
	D             500
	M             1000
	通常情况下，罗马数字中小的数字在大的数字的右边
	I可以放在V(5) 和X(10) 的左边，来表示 4 和 9。
	X可以放在L(50) 和C(100) 的左边，来表示 40 和90。
	C可以放在D(500) 和M(1000) 的左边，来表示400 和900。




*/

func romanToInt(s string) int {
	m := map[byte]int{
		'I':1,
		'V':5,
		'X':10,
		'L':50,
		'C':100,
		'D':500,
		'M':1000,
	}
	var sum int
	var pre int
	for i := len(s) - 1 ; i >= 0 ; i-- {
		cur := m[s[i]]

		if cur < pre {
			sum -= cur
		} else {
			sum += cur
		}
		pre = cur
	}
	return sum

}

func main()  {
	fmt.Println(romanToInt("IV"))
}
