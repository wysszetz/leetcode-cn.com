package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	sLen := len(s)
	if sLen == 0 {
		return 0
	}

	var lastLen int
	for i := sLen - 1; i >= 0; i-- {
		if s[i] != ' ' {
			lastLen++
		}

		if lastLen != 0 && s[i] == ' ' {
			break
		}
	}
	return lastLen
}

func main() {
	s := "Hello World"
	res := lengthOfLastWord(s)
	fmt.Println(res)
}
