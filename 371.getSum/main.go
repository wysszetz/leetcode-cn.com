//不使用运算符 + 和 - ​​​​​​​，计算两整数 ​​​​​​​a 、b ​​​​​​​之和。
//示例 1:
//输入: a = 1, b = 2 输出: 3
//示例 2:
//输入: a = -2, b = 3 输出: 1

package main

import (
	"fmt"
)

func getSum(a int, b int) int {
	sum := a ^ b
	carry := (a & b) << 1
	if carry != 0 {
		return getSum(sum, carry)
	}
	return sum
}

func main() {
	var (
		a = -1
		b = 4
	)
	res := getSum(a, b)
	fmt.Println(res)
}
