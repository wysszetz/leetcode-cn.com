package main

import "fmt"

func moveZeroes(nums []int) []int {
	var (
		newNums []int
		zeroes  int
	)

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroes++
		} else {
			newNums = append(newNums, nums[i])
		}
	}

	for zeroes > 0 {
		newNums = append(newNums, 0)
		zeroes--
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = newNums[i]
	}

	return nums
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	res := moveZeroes(nums)
	fmt.Println(res)
}
