package main

import "fmt"

func init() {
}

func main() {
	fmt.Println()
}

type neighborSum struct {
	adjacentSum []int
	diagonalSum []int
}

func Constructor(grid [][]int) neighborSum {
	rowLen := len(grid)
	colLen := len(grid[0])
	adjacentSum := make([]int, rowLen*colLen)
	diagonalSum := make([]int, rowLen*colLen)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			adjacentSumValue := 0
			if i-1 >= 0 {
				adjacentSumValue += grid[i-1][j]
			}
			if i+1 < rowLen {
				adjacentSumValue += grid[i+1][j]
			}
			if j-1 >= 0 {
				adjacentSumValue += grid[i][j-1]
			}
			if j+1 < colLen {
				adjacentSumValue += grid[i][j+1]
			}
			adjacentSum[grid[i][j]] = adjacentSumValue

			diagonalSumValue := 0
			if i-1 >= 0 && j-1 >= 0 {
				diagonalSumValue += grid[i-1][j-1]
			}
			if i-1 >= 0 && j+1 < colLen {
				diagonalSumValue += grid[i-1][j+1]
			}
			if i+1 < rowLen && j-1 >= 0 {
				diagonalSumValue += grid[i+1][j-1]
			}
			if i+1 < rowLen && j+1 < colLen {
				diagonalSumValue += grid[i+1][j+1]
			}
			diagonalSum[grid[i][j]] = diagonalSumValue
		}
	}
	return neighborSum{adjacentSum: adjacentSum, diagonalSum: diagonalSum}
}

func (this *neighborSum) AdjacentSum(value int) int {
	return this.adjacentSum[value]
}

func (this *neighborSum) DiagonalSum(value int) int {
	return this.diagonalSum[value]
}

/**
 * Your neighborSum object will be instantiated and called as such:
 * obj := Constructor(grid);
 * param_1 := obj.AdjacentSum(value);
 * param_2 := obj.DiagonalSum(value);
 */
