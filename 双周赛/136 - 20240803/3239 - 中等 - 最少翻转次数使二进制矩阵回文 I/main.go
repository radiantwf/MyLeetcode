package main

import "fmt"

func init() {
}

func main() {
	// grid = [[1,0,0],[0,0,0],[0,0,1]] 2
	fmt.Println(minFlips([][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 1}}))
}

func minFlips(grid [][]int) int {
	rowCounter, colCounter := 0, 0
	rowLen, colLen := len(grid), len(grid[0])
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if i < rowLen/2 && grid[i][j] != grid[rowLen-1-i][j] {
				rowCounter++
			}
			if j < colLen/2 && grid[i][j] != grid[i][colLen-1-j] {
				colCounter++
			}
		}
	}
	return min(rowCounter, colCounter)
}
