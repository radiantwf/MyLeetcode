package main

import "fmt"

func main() {
	// grid = [[10,9,0,0,15],[7,1,0,8,0],[5,20,0,11,0],[0,0,0,1,2],[8,12,1,10,3]]
	// 输出：94
	grid := [][]int{
		{10, 9, 0, 0, 15},
		{7, 1, 0, 8, 0},
		{5, 20, 0, 11, 0},
		{0, 0, 0, 1, 2},
		{8, 12, 1, 10, 3},
	}
	fmt.Println(maximumScore(grid))
}
func maximumScore(grid [][]int) int64 {
	return 0
}
