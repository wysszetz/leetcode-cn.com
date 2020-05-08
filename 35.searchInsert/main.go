package main

import "fmt"

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	if target < nums[0] {
		return 0
	}

	numsLen := len(nums)

	if target > nums[numsLen-1] {
		return numsLen
	}

	lower, high := 0, numsLen-1

	for lower <= high {
		middle := (lower + high) / 2
		if nums[middle] < target {
			lower = middle + 1
		} else if nums[middle] > target {
			high = middle - 1
		} else {
			return middle
		}
	}
	return high + 1
}

func main() {
	nums := []int{1, 3, 4, 6}
	target := 5
	res := searchInsert(nums, target)
	fmt.Println(res)
}
