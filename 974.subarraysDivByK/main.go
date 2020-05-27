package main

import "fmt"

func subarraysDivByK(A []int, K int) int {
	var record = map[int]int{0: 1}
	sum, answer := 0, 0

	for _, val := range A {
		sum += val
		fmt.Print(sum, "\t")
		mod := (sum%K + K) % K
		fmt.Print(mod, "\t")
		answer += record[mod]
		fmt.Print(answer, "\t")
		record[mod]++
		fmt.Println(record)
	}
	return answer
}

func main() {
	A := []int{4, 5, 0, -2, -3, 1}
	K := 5

	res := subarraysDivByK(A, K)
	fmt.Println(res)
}
