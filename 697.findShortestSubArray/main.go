package main

import (
	"fmt"
)

func findShortestSubArray(nums []int) int {
	var (
		left  = make(map[int]int)
		right = make(map[int]int)
		count = make(map[int]int)
	)

	for i := 0; i < len(nums); i++ {
		tmp := nums[i]
		_, ok := left[tmp]
		if !ok {
			left[tmp] = i
		}
		right[tmp] = i
		count[nums[i]]++
	}

	answer := len(nums)
	degree := 0
	for _, v := range count {
		if v >= degree {
			degree = v
		}
	}

	for k, v := range count {
		if degree == v {
			length := right[k] - left[k] + 1
			if length <= answer {
				answer = length
			}
		}
	}

	return answer
}
func main() {
	nums := []int{1, 2, 2, 3, 1, 4, 2}
	res := findShortestSubArray(nums)
	fmt.Println(res)
}
