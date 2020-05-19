package main

import "fmt"

func validPalindrome(s string) bool {
	var low, high = 0, len(s) - 1

	for low <= high {
		if s[low] == s[high] {
			low++
			high--
		} else {
			var flag1, flag2 = true, true
			i, j := low+1, high
			for i < j {
				if s[i] != s[j] {
					flag1 = false
					break
				} else {
					i++
					j--
				}
			}

			i, j = low, high-1
			for i < j {
				if s[i] != s[j] {
					flag2 = false
					break
				} else {
					i++
					j--
				}
			}

			return flag1 || flag2
		}
	}
	return true
}

func main() {
	s := "abca"
	res := validPalindrome(s)
	fmt.Println(res)
}
