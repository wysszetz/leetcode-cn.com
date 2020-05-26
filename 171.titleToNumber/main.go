package main

import "fmt"

func titleToNumber(s string) int {
	var answer int
	for i := 0; i < len(s); i++ {
		num := s[i] - 'A' + 1
		answer = answer*26 + int(num)
	}

	return answer
}

func main() {
	s := "ZY"
	res := titleToNumber(s)
	fmt.Println(res)
}
