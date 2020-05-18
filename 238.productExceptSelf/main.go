package main

import "fmt"

func productExceptSelf(nums []int) []int {
	var output, left, right = make([]int, len(nums)), make([]int, len(nums)), make([]int, len(nums))

	left[0], right[len(nums)-1] = 1, 1

	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
		right[len(nums)-i-1] = right[len(nums)-i] * nums[len(nums)-i]
	}
	fmt.Println(left, right)

	for i := 0; i < len(output); i++ {
		output[i] = left[i] * right[i]
	}

	return output
}

func main() {
	nums := []int{1, 2, 3, 4}
	res := productExceptSelf(nums)
	fmt.Println(res)
}
