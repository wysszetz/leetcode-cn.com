package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	var (
		result []int
	)

	var numsLen = len(nums)

	for i := 0; i <= numsLen-k; i++ {
		tmp := append([]int{}, nums[i:i+k]...)
		tmpMax := tmp[0]
		for _, v := range tmp {
			if tmpMax <= v {
				tmpMax = v
			}
		}
		result = append(result, tmpMax)
	}

	return result
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	res := maxSlidingWindow(nums, k)
	fmt.Println(res)
}
