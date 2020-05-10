package main

import (
	"fmt"
	"math"
)

func longestPalindrome(s string) string {
	sLen := len(s)
	if sLen <= 1 {
		return s
	}
	start, end := 0, 0
	for i := 0; i < sLen; i++ {
		maxLen := int(math.Max(float64(max(s, i, i)), float64(max(s, i, i+1))))
		if maxLen > (end - start + 1) {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}
	return s[start : end+1]
}

func max(s string, left, right int) int {
	L, R := left, right
	for (L >= 0 && R < len(s)) && s[L] == s[R] {
		R++
		L--
	}
	return R - L - 1
}

func main() {
	str := "cbbd"
	res := longestPalindrome(str)
	fmt.Println(res)
}
