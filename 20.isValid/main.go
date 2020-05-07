//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//有效字符串需满足：
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//注意空字符串可被认为是有效字符串。
//示例 1:输入: '()'输出: true
//示例 2:输入: '()[]{}'输出: true
//示例 3:输入: '(]'输出: false
//示例 4:输入: '([)]'输出: false
//示例 5:输入: '{[]}'输出: true
package main

import "fmt"

func isValid(s string) bool {
	strLen := len(s)
	if strLen%2 != 0 {
		return false
	}
	var (
		isValid bool
		valid   = map[uint8]uint8{')': '(', '}': '{', ']': '['}
		stack   []uint8
	)

	for i := 0; i < strLen; i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return isValid
			}
			if valid[s[i]] != stack[len(stack)-1] {
				return isValid
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	res := isValid("()")
	fmt.Println(res)
}
