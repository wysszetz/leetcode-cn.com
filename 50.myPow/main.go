package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return helper(x, n)
	}
	return 1.0 / helper(x, -n)
}

func helper(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := helper(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}

func main() {
	res := myPow(2.00000, 10)
	fmt.Println(res)
}
