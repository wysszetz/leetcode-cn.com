package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	var count = make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}
	for k, v := range count {
		if v == 1 {
			return k
		}
	}
	return 0
}

func main() {
	nums := []int{2, 2, 1}
	res := singleNumber(nums)
	fmt.Println(res)
}
