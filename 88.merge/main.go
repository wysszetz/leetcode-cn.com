package main

import "fmt"
// 双指针
func merge(nums1 []int, m int, nums2 []int, n int) []int {
	if len(nums1) == 0 {
		return nums2
	}
	if len(nums2) == 0 {
		return nums1
	}

	var (
		mLen = m - 1
		nLen = n - 1
		len  = m + n - 1
	)

	for nLen >= 0 {
		if mLen >= 0 && nums1[mLen] > nums2[nLen] {
			nums1[len] = nums1[mLen]
			mLen--
		} else {
			nums1[len] = nums2[nLen]
			nLen--
		}
		len--
	}
	return nums1
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	res := merge(nums1, m, nums2, n)
	fmt.Println(res)
}
