package main

import "fmt"

func findErrorNums(nums []int) []int {
	var (
		result     []int
		docker     = make(map[int]int, len(nums))
		miss, copy = -1, -1
	)

	for i := 0; i < len(nums); i++ {
		docker[nums[i]]++
	}

	for i := 0; i < len(nums); i++ {
		if _, ok := docker[i+1]; !ok {
			miss = i + 1
		}

		if docker[i+1] > 1 {
			copy = i + 1
		}
	}

	result = append(result, copy)
	result = append(result, miss)
	return result
}

func main() {
	nums := []int{3, 2, 3, 4, 6, 5}
	res := findErrorNums(nums)
	fmt.Println(res)
}
