package main

import "fmt"

func subarraySum(nums []int, k int) int {
	var count int
	for i := 0; i < len(nums); i++ {
		var sum int
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}
	return count
}
func main() {
	nums := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	k := 0
	res := subarraySum(nums, k)
	fmt.Println(res)
}
