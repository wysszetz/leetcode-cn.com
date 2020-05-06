//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//注意：给定 n 是一个正整数。
//
//示例 1：
//
//输入： 2
//输出： 2
//解释： 有两种方法可以爬到楼顶。
//1.  1 阶 + 1 阶
//2.  2 阶
//示例 2：
//
//输入： 3
//输出： 3
//解释： 有三种方法可以爬到楼顶。
//1.  1 阶 + 1 阶 + 1 阶
//2.  1 阶 + 2 阶
//3.  2 阶 + 1 阶

package main

import "fmt"

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	stairs := make(map[int]int, 2)
	stairs[1] = 1
	stairs[2] = 2

	for i := 3; i <= n; i++ {
		stairs[i] = stairs[i-1] + stairs[i-2]
	}
	return stairs[n]
}

func main() {
	res := climbStairs(3)
	fmt.Println(res)
}
