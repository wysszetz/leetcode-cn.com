package main

import "fmt"

func maximumProduct(nums []int) int {
	nums = quickSort(nums)
	var (
		tmpNums    int
		tmpMaxNums = 1
	)
	if nums[0] < 0 && nums[1] < 0 {
		tmpNums = nums[0] * nums[1]
	}

	for i := len(nums) - 3; i < len(nums); i++ {
		tmpMaxNums = tmpMaxNums * nums[i]
	}
	return max(tmpMaxNums, tmpNums*nums[len(nums)-1])
}

func quickSort(nums []int) []int {
	numsLen := len(nums)
	if numsLen <= 1 {
		return nums
	}

	var left, right = []int{}, []int{}
	for i := 1; i < numsLen; i++ {
		if nums[i] < nums[0] {
			left = append(left, nums[i])
		} else {
			right = append(right, nums[i])
		}
	}
	left = quickSort(left)
	right = quickSort(right)

	newNums := append([]int{}, left...)
	newNums = append(newNums, nums[0])
	newNums = append(newNums, right...)
	return newNums
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{1, 2, 3, 4, -10, -11}
	res := maximumProduct(nums)
	fmt.Println(res)
}
