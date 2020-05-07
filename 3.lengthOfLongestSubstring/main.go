//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//示例 1:
//输入: "abcabcbb" 输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//输入: "bbbbb" 输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//输入: "pwwkew" 输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

package main

import (
	"fmt"
	"math"
)

//滑动窗口
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 { // 当输入字符串为空时，返回0
		return 0
	}
	var (
		strMap = make(map[byte]int)
		strLen = len(s)
		index  = -1
		answer = 0
	)

	for i := 0; i < strLen; i++ {
		if i != 0 {
			delete(strMap, s[i-1])
		}
		for index+1 < strLen && strMap[s[index+1]] == 0 {
			strMap[s[index+1]]++
			index++
		}
		answer = int(math.Max(float64(answer), float64(index-i+1)))
	}
	return answer
}

func main() {
	res := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(res)
}
