package main

import "fmt"

func bitwiseComplement(N int) int {
	if N == 0 {
		return 1
	}
	i := 1
	for i <= N {
		i <<= 1
		fmt.Println(i)
	}
	return i - N - 1
}

func main() {
	N := 7
	res := bitwiseComplement(N)
	fmt.Println(res)
}
