package main

import (
	"fmt"
)

func maxProduct(nums []int) int {
	var maxF, minF, answer = nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max(mx*nums[i], max(mn*nums[i], nums[i]))
		minF = min(mn*nums[i], min(mx*nums[i], nums[i]))
		answer = max(maxF, answer)
	}
	return answer
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func main() {
	nums := []int{2, 3, -2, 4}
	res := maxProduct(nums)
	fmt.Println(res)
}
