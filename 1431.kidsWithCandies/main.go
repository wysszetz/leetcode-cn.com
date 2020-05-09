package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var (
		result     []bool
		maxCandies int
	)

	for _, v := range candies {
		if v >= maxCandies {
			maxCandies = v
		}
	}
	minCandies := maxCandies - extraCandies
	for _, v := range candies {
		if v < minCandies {
			result = append(result, false)
		} else {
			result = append(result, true)
		}
	}
	return result
}
func main() {
	//candies := []int{2, 3, 5, 1, 3}
	//extraCandies := 3
	candies := []int{4, 2, 1, 1, 2}
	extraCandies := 1
	res := kidsWithCandies(candies, extraCandies)
	fmt.Println(res)
}
