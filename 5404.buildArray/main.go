package main

import "fmt"

func buildArray(target []int, n int) []string {
	var list []int
	for i := 1; i <= n; i++ {
		list = append(list, i)
	}
	targetLen := len(target)
	fmt.Println(list, target, targetLen)

	var result []string
	var lastList int
	for i := 0; i < targetLen; i++ {
		var tmp = target[i]
		for j := lastList; j < n; j++ {
			if list[j] < tmp {
				result = append(result, "Push")
				result = append(result, "Pop")
			} else if list[j] == tmp {
				result = append(result, "Push")
				lastList = j + 1
				break
			} else {
				break
			}
		}
	}

	return result
}

func main() {
	target := []int{1, 2, 3}
	n := 3
	res := buildArray(target, n)
	fmt.Println(res)
}
