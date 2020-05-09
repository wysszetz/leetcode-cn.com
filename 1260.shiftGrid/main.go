package main

import "fmt"

func shiftGrid(grid [][]int, k int) [][]int {
	m := len(grid)
	n := len(grid[0])
	if m*n <= k && k%(m*n) == 0 {
		return grid
	}
	mod := k % (m * n)
	var newGrid []int

	for _, v := range grid {
		newGrid = append(newGrid, v...)
	}
	newGridLen := len(newGrid)
	tmp1 := append([]int{}, newGrid[:newGridLen-mod]...)
	tmp2 := append([]int{}, newGrid[newGridLen-mod:]...)
	tmp2 = append(tmp2, tmp1...)
	newGrid = tmp2

	var newGridResult [][]int

	for i := 0; i < m; i++ {
		tmp := append([]int{}, newGrid[i*n:i*n+n]...)
		newGridResult = append(newGridResult, tmp)
	}

	return newGridResult
}
func main() {
	//gird := [][]int{{3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}, {12, 0, 21, 13}}
	//k := 4

	gird := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	k := 1
	res := shiftGrid(gird, k)
	fmt.Println(res)
}
