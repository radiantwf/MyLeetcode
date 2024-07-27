package main

import (
	"fmt"
)

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

// 解题思路：
// x：表示涂黑 1：表示不涂黑，计算该格的分值 0：表示不涂黑，不计算分值
// 先观察 涂黑的格子情况，可以发现
// impossible	possible
// x x x		x 1 x
// x x x		x 1 x
// x 1 x		x 1 x
// x 1 0		x 1 0
// 0 0 0		0 0 0
// 对比左右两个图，可以发现，左图的涂黑方式（凹字式）在任何情况下，不会比右图（隔列式）更高（每格分值>=0）,所以计算时可以不考虑左图情况

// 可能的情况有：
// x x x    x x x	x x x	x 1 x	x 1 x
// x x x	x x x	x x x	x 1 x	x 1 x
// 1 x x	x x 1	x x 1	x 1 x	x 1 x
// 1 1 x	x 1 0	1 x 1	0 1 x	x 1 0
// 0 0 0	0 0 0	0 0 0	0 0 0	0 0 0
// 递增式、递减式、凸字式（先递增再递减）、隔列式（4、5图） 4种情况

// 综上所述：
// 当前列涂黑行数较前一列增加时：有效
// 当前列涂黑行数较前一列减少时：若前列与前前列是递增趋势，则无效，此列应该不涂黑，否则有效

func maximumScore(grid [][]int) int64 {
	rowCount := len(grid)
	colCount := len(grid[0])

	// 计算前缀和，为了方便计算不涂黑任何一个格子，扩展一行
	prefixSum := make([][]int64, rowCount+1)
	prefixSum[0] = make([]int64, colCount)
	for row := 0; row < len(grid); row++ {
		prefixSum[row+1] = make([]int64, colCount)
		for col := 0; col < len(grid[row]); col++ {
			if row > 0 {
				prefixSum[row+1][col] = prefixSum[row][col] + int64(grid[row][col])
			} else {
				prefixSum[row+1][col] = int64(grid[row][col])
			}
		}
	}
	// 为了变量数组时初始化方便，dp数组的第一维度为列数;涂黑格子时，可以选择不涂黑，所以多一行
	// [2] 维度0表示递减式，维度1表示递增式
	dp := make([][][2]*int64, colCount)
	dp[0] = make([][2]*int64, rowCount+1)
	var getDp func(col, leftColBlackRow, order int) int64
	getDp = func(col, leftColBlackRow, order int) int64 {
		if col >= colCount {
			return 0
		}
		if dp[col] == nil {
			dp[col] = make([][2]*int64, rowCount+1)
		}
		if dp[col][leftColBlackRow][order] != nil {
			return *dp[col][leftColBlackRow][order]
		}
		source := int64(0)
		for blackRow := 0; blackRow < rowCount+1; blackRow++ {
			if col == 0 {
				source = max(source, getDp(col+1, blackRow, 1))
				continue
			}
			if blackRow == leftColBlackRow {
				// 相等，不计算分值，直接处理下一列
				source = max(source, getDp(col+1, blackRow, 1))
			} else if blackRow < leftColBlackRow {
				// 递减式
				source = max(source, getDp(col+1, blackRow, 0)+prefixSum[leftColBlackRow][col]-prefixSum[blackRow][col])
			} else if blackRow > leftColBlackRow {
				// 递增式
				if order == 0 {
					if leftColBlackRow != 0 {
						continue
					}
					source = max(source, getDp(col+1, blackRow, 1))
				} else {
					if col == 0 {
						source = max(source, getDp(col+1, blackRow, 1))
					} else {
						source = max(source, getDp(col+1, blackRow, 1)+prefixSum[blackRow][col-1]-prefixSum[leftColBlackRow][col-1])
					}
				}
			}
		}
		value := source
		dp[col][leftColBlackRow][order] = &value
		return value
	}
	source := getDp(0, 0, 1)
	return source
}
