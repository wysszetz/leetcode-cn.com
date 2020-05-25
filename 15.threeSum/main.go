package main

import "fmt"

func threeSum(nums []int) [][]int {
	nums = quickSort(nums)

	var (
		resultNums [][]int
		numsLen    = len(nums)
	)

	if numsLen < 3 {
		return resultNums
	}

	for i := 0; i < numsLen; i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := numsLen - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				resultNums = append(resultNums, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return resultNums
}

func quickSort(nums []int) []int {
	numsLen := len(nums)
	if numsLen <= 1 {
		return nums
	}

	var left, right []int

	for i := 1; i < numsLen; i++ {
		if nums[0] > nums[i] {
			left = append(left, nums[i])
		} else {
			right = append(right, nums[i])
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	tmp := append([]int{}, left...)
	tmp = append(tmp, nums[0])
	tmp = append(tmp, right...)
	return tmp

}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	fmt.Println(res)
}
