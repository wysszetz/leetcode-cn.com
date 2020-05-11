package main

import (
	"fmt"
	"math"
)

func rob(nums []int) int {
	var (
		prex  int
		curry int
	)

	for _, v := range nums {
		var tmp = curry
		curry = int(math.Max(float64(prex+v), float64(curry)))
		prex = tmp
	}
	return curry
}

func main() {
	nums := []int{2, 7, 9, 100, 1}
	res := rob(nums)
	fmt.Println(res)
}
