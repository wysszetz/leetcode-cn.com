package main

import "fmt"

func findDuplicate(nums []int) int {
	n := len(nums)
	l, r := 1, n-1
	ans := -1
	for l <= r {
		mid := (l + r) >> 1
		cnt := 0
		for i := 0; i < n; i++ {
			if nums[i] <= mid {
				cnt++
			}
		}
		fmt.Println(mid, cnt)
		if cnt <= mid {
			l = mid + 1
		} else {
			r = mid - 1
			ans = mid
		}
	}
	return ans
}

func main() {
	nums := []int{1, 3, 4, 2, 2}
	res := findDuplicate(nums)
	fmt.Println(res)
}
