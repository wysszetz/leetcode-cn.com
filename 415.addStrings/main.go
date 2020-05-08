package main

import (
	"fmt"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	num1Len, num2Len := len(num1)-1, len(num2)-1

	var (
		newNumStr = ""
		flag      int
	)

	for num1Len >= 0 || num2Len >= 0 {
		num1Num, num2Num := 0, 0
		if num1Len >= 0 {
			num1Num, _ = strconv.Atoi(string(num1[num1Len]))
		}
		if num2Len >= 0 {
			num2Num, _ = strconv.Atoi(string(num2[num2Len]))
		}
		sum := num1Num + num2Num + flag
		if sum > 9 {
			flag = sum / 10
			sum -= 10
		} else {
			flag = 0
		}
		newNumStr = strconv.Itoa(sum) + newNumStr
		num1Len--
		num2Len--
	}

	if flag > 0 {
		newNumStr = strconv.Itoa(flag) + newNumStr
	}
	return newNumStr
}

func main() {
	num1, num2 := "121", "999"
	res := addStrings(num1, num2)
	fmt.Println(res)
}
