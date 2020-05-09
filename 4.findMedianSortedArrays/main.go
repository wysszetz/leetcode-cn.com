package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var newNums []int
	newNums = append(newNums, nums1...)
	newNums = append(newNums, nums2...)
	for i := 0; i < len(newNums)-1; i++ {
		var p = i
		for j := i + 1; j < len(newNums); j++ {
			if newNums[p] > newNums[j] {
				p = j
			}
		}
		if p != i {
			tmp := newNums[p]
			newNums[p] = newNums[i]
			newNums[i] = tmp
		}
	}
	newNumsLen := len(newNums)

	if newNumsLen%2 == 1 {
		middle := int(math.Ceil(float64(newNumsLen / 2)))
		return float64(newNums[middle])
	} else {
		middle := newNumsLen/2 + 1
		return float64(newNums[middle-1]+newNums[middle-2]) / 2
	}
}
func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2, 9, 8, 10}
	res := findMedianSortedArrays(nums1, nums2)
	fmt.Println(res)
}
