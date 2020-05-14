package main

import (
	"fmt"
)

func singleNumber(nums []int) []int {
	var count = make(map[int]int, len(nums))
	var result []int
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}
	for k, v := range count {
		if v == 1 {
			result = append(result, k)
		}
	}
	return result
}

func main() {
	nums := []int{1, 2, 1, 3, 2, 5}
	res := singleNumber(nums)
	fmt.Println(res)
}
