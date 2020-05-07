//编写一个函数来查找字符串数组中的最长公共前缀。
//如果不存在公共前缀，返回空字符串 ""。
//示例 1:
//输入: ["flower","flow","flight"]输出: "fl"
//示例 2:
//输入: ["dog","racecar","car"] 输出: ""
//解释: 输入不存在公共前缀。
//说明:所有输入只包含小写字母 a-z 。

package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	strsLen := len(strs)
	if strsLen == 0 {
		return ""
	}
	var (
		minStrLen = len(strs[0])
		minStr    = strs[0]
		answer    = ""
	)
	for _, val := range strs {
		if minStrLen >= len(val) {
			minStrLen = len(val)
			minStr = val
		}
	}

	for i := 0; i < minStrLen; i++ {
		offset := true
		for _, val := range strs {
			if val[i] != minStr[i] {
				offset = false
			}
		}
		if offset {
			answer += string(minStr[i])
		} else {
			break
		}
	}
	return answer
}

func main() {
	//strs := []string{"flower", "flow", "flight"}
	strs := []string{"dog","racecar","car"}
	res := longestCommonPrefix(strs)
	fmt.Println(res)
}
