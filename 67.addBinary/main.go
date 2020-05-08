//给你两个二进制字符串，返回它们的和（用二进制表示）。
//输入为 非空 字符串且只包含数字 1 和 0。
//示例 1:
//输入: a = "11", b = "1" 输出: "100"
//示例 2:
//输入: a = "1010", b = "1011" 输出: "10101"
//提示：
//每个字符串仅由字符 '0' 或 '1' 组成。
//1 <= a.length, b.length <= 10^4
//字符串如果不是 "0" ，就都不含前导零。

package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	var (
		newBinary = ""
		carry     int
		flag      int
	)
	aLen, bLen := len(a)-1, len(b)-1

	for aLen >= 0 || bLen >= 0 {
		aInt, bInt := 0, 0
		if aLen >= 0 {
			aInt = int(a[aLen] - '0')
		}
		if bLen >= 0 {
			bInt = int(b[bLen] - '0')
		}
		carry = aInt + bInt + flag
		flag = 0
		if carry >= 2 {
			flag = 1
			carry = carry - 2
		}

		newBinary = strconv.Itoa(carry) + newBinary

		aLen--
		bLen--
	}
	if flag == 1 {
		newBinary = strconv.Itoa(flag) + newBinary
	}
	return newBinary
}

func main() {
	a := "11"
	b := "1"
	res := addBinary(a, b)
	fmt.Println(res)
}
