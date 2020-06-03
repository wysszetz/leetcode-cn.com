package main

import "fmt"

func sumNums(n int) int {
	return ((1 + n) * n) / 2
}

func main() {
	n := 9
	res := sumNums(n)
	fmt.Println(res)
}
