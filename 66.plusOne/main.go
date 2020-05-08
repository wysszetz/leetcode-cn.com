package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	digitsLen := len(digits)
	var (
		newDigits []int
		flag      int
	)
	for i := digitsLen - 1; i >= 0; i-- {
		var sum int
		if i == digitsLen-1 {
			sum = digits[i] + 1
		} else {
			sum = flag + digits[i]
		}
		flag = sum / 10
		newDigits = append(newDigits, sum%10)
	}
	if flag > 0 {
		newDigits = append(newDigits, flag)
	}

	for i, j := 0, len(newDigits)-1; i < j; i, j = i+1, j-1 {
		newDigits[i], newDigits[j] = newDigits[j], newDigits[i]
	}

	return newDigits
}

func main() {
	digits := []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}
	res := plusOne(digits)
	fmt.Println(res)
}
