package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	isValid := func(s rune) bool {
		return unicode.IsLetter(s) || unicode.IsDigit(s)
	}

	var low, high = 0, len(s) - 1

	for low <= high {
		var tmp1, tmp2 = rune(s[low]), rune(s[high])
		if !isValid(tmp1) && !isValid(tmp2) {
			low++
			high--
		} else if !isValid(tmp1) {
			low++
		} else if !isValid(tmp2) {
			high--
		} else if unicode.ToUpper(tmp1) != unicode.ToUpper(tmp2) {
			return false
		} else {
			low++
			high--
		}
	}
	return true
}
func main() {
	s := "A man, a plan, a canal: Panama"
	res := isPalindrome(s)
	fmt.Println(res)
}
