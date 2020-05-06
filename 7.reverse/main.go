//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
//
//示例 1:
//输入: 123
//输出: 321
// 示例 2:
//输入: -123
//输出: -321
//示例 3:isPalindrome
//输入: 120
//输出: 21

package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var (
		tmp      []int
		negative bool
	)
	if x < 0 {
		negative = true
		x = x * -1
	} else {
		negative = false
	}
	for x != 0 {
		mod := x % 10
		x = (x - mod) / 10
		tmp = append(tmp, mod)
	}
	tmpLen := len(tmp) - 1
	newSum := 0

	for _, v := range tmp {
		tmp10 := int(math.Pow10(tmpLen))
		tmpNum := v * tmp10
		newSum = newSum + tmpNum
		tmpLen--
	}

	if negative {
		newSum = newSum * -1
		if newSum < int(math.Pow(2, 31))*-1 {
			return 0
		}
	} else {
		if newSum > int(math.Pow(2, 31)-1) {
			return 0
		}
	}

	return newSum
}

func main() {
	res := reverse(200)
	fmt.Println(res)
}
