package main

import (
	"fmt"
)

func sortedSquares(A []int) []int {
	for k, v := range A {
		A[k] = v * v
	}
	for i := 0; i < len(A)-1; i++ {
		var p = i
		for j := i + 1; j < len(A); j++ {
			if A[p] > A[j] {
				p = j
			}
		}
		if p != i {
			tmp := A[p]
			A[p] = A[i]
			A[i] = tmp
		}
	}

	return A
}

func main() {
	A := []int{-4, -1, 0, 3, 10}
	res := sortedSquares(A)
	fmt.Println(res)
}
