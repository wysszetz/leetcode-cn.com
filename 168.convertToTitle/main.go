package main

import "fmt"

func convertToTitle(n int) string {
	var title string
	for n > 0 {
		mod := n % 26
		if mod == 0 {
			mod = 26
			n -= 1
		}
		n = n / 26
		title = string(mod+64) + title
	}
	return title
}

func main() {
	n := 701
	res := convertToTitle(n)
	fmt.Println(res)
}
