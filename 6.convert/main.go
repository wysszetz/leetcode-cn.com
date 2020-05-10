package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var (
		sLen      = len(s)
		newString string
		cycleLen  = 2*numRows - 2
	)

	for i := 0; i < numRows; i++ {
		for j := 0; j+i < sLen; j += cycleLen {
			newString += string(s[j+i])
			if i != 0 && i != numRows-1 && j+cycleLen-i < sLen {
				newString += string(s[j+cycleLen-i])
			}
		}
	}
	return newString
}

func main() {
	s := "LEETCODEISHIRING"
	res := convert(s, 3)
	fmt.Println(res)
}
