package main

import (
	"fmt"
	"strconv"
)

func multiply(num1 string, num2 string) string {
	num1Len, num2Len := len(num1)-1, len(num2)-1

	if num2Len > num1Len {
		return multiply(num2, num1)
	}

	var (
		newNum    []string
		newNumStr = ""
	)

	for i := num2Len; i >= 0; i-- {
		num2Num, _ := strconv.Atoi(string(num2[i]))
		var (
			flag   int
			tmpStr string
		)
		for j := num1Len; j >= 0; j-- {
			num1Num, _ := strconv.Atoi(string(num1[j]))
			pro := num2Num*num1Num + flag
			if pro > 9 {
				flag = pro / 10
				pro = pro % 10
			} else {
				flag = 0
			}
			tmpStr = strconv.Itoa(pro) + tmpStr
		}
		if flag > 0 {
			tmpStr = strconv.Itoa(flag) + tmpStr
		}
		pow := num2Len - i
		if pow >= 1 {
			for k := 1; k <= num2Len-i; k++ {
				tmpStr += "0"
			}
		}
		newNum = append(newNum, tmpStr)
	}
	if len(newNum) == 1 {
		newNumStr = newNum[0]
	} else {
		newNumStr = newNum[0]
		for i := 1; i < len(newNum); i++ {
			newNumStr = addStrings(newNumStr, newNum[i])
		}
	}

	var zero int
	for _, v := range newNumStr {
		if v == '0' {
			zero++
		}
	}
	if zero == len(newNumStr) {
		return "0"
	}
	return newNumStr
}

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
	num1, num2 := "9133", "0"
	res := multiply(num1, num2)
	fmt.Println(res)
}
