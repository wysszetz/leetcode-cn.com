package main

import "fmt"

func missingNumber(nums []int) int {
	nums = QuickSort(nums)
	if nums[0] != 0 {
		return 0
	}

	if nums[len(nums)-1] != len(nums) {
		return len(nums)
	}

	for i := 1; i < len(nums); i++ {
		tmp := nums[i-1] + 1
		if tmp != nums[i] {
			return tmp
		}
	}
	return -1
}

func QuickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	var left, right []int

	for i := 1; i < len(nums); i++ {
		if nums[0] > nums[i] {
			left = append(left, nums[i])
		} else {
			right = append(right, nums[i])
		}
	}

	left = QuickSort(left)
	right = QuickSort(right)

	left = append(left, nums[0])
	return append(left, right...)
}

func main() {
	nums := []int{0, 2, 3}
	res := missingNumber(nums)
	fmt.Println(res)

}
