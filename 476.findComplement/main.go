package main

import "fmt"

func findComplement(num int) int {
	i := 1
	for i <= num {
		i = i << 1
	}
	return i - num - 1
}

func main() {
	num := 5
	res := findComplement(num)
	fmt.Println(res)
}
