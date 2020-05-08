package main

import (
	"fmt"
)

func addToArrayForm(A []int, K int) []int {
	var (
		aLen   = len(A) - 1
		newInt []int
		flag   int
	)
	for aLen >= 0 || K > 0 {
		aInt, kMod := 0, 0
		if aLen >= 0 {
			aInt = A[aLen]
		}
		if K > 0 {
			kMod = K % 10
		}
		sum := kMod + aInt + flag
		if sum > 9 {
			flag = sum / 10
			sum -= 10
		} else {
			flag = 0
		}
		newInt = append(newInt, sum)
		K = (K - kMod) / 10
		aLen--
	}
	if flag > 0 {
		newInt = append(newInt, flag)
	}

	for i, j := 0, len(newInt)-1; i < j; i, j = i+1, j-1 {
		newInt[i], newInt[j] = newInt[j], newInt[i]
	}

	return newInt
}

func main() {
	a := []int{2, 7, 4}
	k := 181
	res := addToArrayForm(a, k)
	fmt.Println(res)
}
