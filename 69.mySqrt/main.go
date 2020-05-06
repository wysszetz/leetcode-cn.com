//实现 int sqrt(int x) 函数。
//
//计算并返回 x 的平方根，其中 x 是非负整数。
//
//由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
//
//示例 1:
//
//输入: 4
//输出: 2
//示例 2:
//
//输入: 8
//输出: 2
//说明: 8 的平方根是 2.82842...,
//     由于返回类型是整数，小数部分将被舍去。
package main

import "fmt"

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	var (
		left   = 2
		middle int
		right  = x / 2
	)

	for left <= right {
		middle = left + (right-left)/2
		num := middle * middle
		if num > x {
			right = middle - 1
		} else if num < x {
			left = middle + 1
		} else {
			return middle
		}
	}
	return right
}

func main() {
	res := mySqrt(8)
	fmt.Println(res)
}
