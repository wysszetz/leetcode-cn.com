package main

import (
	"fmt"
	"math"
)

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	nums1 := append([]int{}, nums[1:]...)
	nums2 := append([]int{}, nums[:len(nums)-1]...)
	nums1Max := helper(nums1)
	nums2Max := helper(nums2)
	if nums1Max >= nums2Max {
		return nums1Max
	} else {
		return nums2Max
	}
}

func helper(nums []int) int {
	var (
		prex  int
		curry int
	)

	for _, v := range nums {
		var tmp = curry
		curry = int(math.Max(float64(prex+v), float64(curry)))
		prex = tmp
	}
	return curry
}

func main() {
	nums := []int{1, 2, 3, 1}
	res := rob(nums)
	fmt.Println(res)
}
