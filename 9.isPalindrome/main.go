package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x >= 0 && x <= 9 {
		return true
	}
	var tmp []int
	for x > 0 {
		mod := x % 10
		x = (x - mod) / 10
		tmp = append(tmp, mod)
	}
	left := 0
	right := len(tmp) - 1
	var isPalindrome bool
	for left <= right {
		if tmp[left] != tmp[right] {
			isPalindrome = false
			break
		} else {
			left++
			right--
			isPalindrome = true
		}
	}

	return isPalindrome
}

func main() {
	fmt.Println(123)
	result := isPalindrome(0)
	fmt.Println(result)
}
