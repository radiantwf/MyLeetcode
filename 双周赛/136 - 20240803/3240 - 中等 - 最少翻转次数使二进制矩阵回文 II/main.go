package main

import (
	"fmt"
)

func init() {
}

func main() {
	// grid = [[0,0,0],[1,1,0],[0,1,1],[0,0,1]]
	// fmt.Println(minFlips([][]int{{0, 0, 0}, {1, 1, 0}, {0, 1, 1}, {0, 0, 1}}))

	fmt.Println(minFlips([][]int{{1}, {1}, {0}, {1}, {1}}))
}

// 解题思路：
// 1. 节点的行列的对称节点为4个点时，无论如何反转，只要形成回文，必然是偶数个1，所以只需要记录最小反转次数即可；
// 1.1 1或3个节点为1，其他为0，反转次数加1
// 1.2 2个节点为1，其他为0，反转次数加2

// 2. 节点的行列的对称节点只有自己时 奇数行列矩阵的中心节点，由于只有1个节点，而且只要这个节点为1，回文矩阵的1的个数一定为奇数，所以必须反转为0；
// 2.1 节点值为1时，反转次数加1

// 3. 节点的行列的对称节点为2个点时 奇数行矩阵的中间行，奇数列矩阵的中间列（不包含奇数行列矩阵的中心节点），只要这类节点1的个数能被4整除，结果即成立；
// 3.1 当1个节点为1，另一个节点为0时，反转次数加1
// 3.2 当所有节点1的个数不能被4整除时，且没有1组需要反转时（只要有1组需要反转，即可调整），反转次数加2

func minFlips(grid [][]int) int {
	count := 0
	midOneFixFlag := false
	midFixNeedFlipFlag := true

	rowLen, colLen := len(grid), len(grid[0])
	midRow := (rowLen + 1) / 2
	midCol := (colLen + 1) / 2

	for i := 0; i < midRow; i++ {
		for j := 0; j < midCol; j++ {
			v := grid[i][j]
			vr := -1
			vc := -1
			vcr := -1
			if rowLen-1-i != i {
				vr = grid[rowLen-1-i][j]
			}
			if colLen-1-j != j {
				vc = grid[i][colLen-1-j]
			}
			if rowLen-1-i != i && colLen-1-j != j {
				vcr = grid[rowLen-1-i][colLen-1-j]
			}
			sum := v + max(vr, 0) + max(vc, 0) + max(vcr, 0)
			if vr != -1 && vc != -1 {
				if vr != v || vc != v || v != vcr {
					if sum == 1 || sum == 3 {
						count++
					} else if sum == 2 {
						count += 2
					}
				}
			} else if vr != -1 {
				if vr != v {
					count++
					midFixNeedFlipFlag = false
				} else if v == 1 {
					midOneFixFlag = !midOneFixFlag
				}
			} else if vc != -1 {
				if vc != v {
					count++
					midFixNeedFlipFlag = false
				} else if v == 1 {
					midOneFixFlag = !midOneFixFlag
				}
			} else if vc == -1 && vr == -1 {
				if v == 1 {
					count++
				}
			}
		}
	}
	if midOneFixFlag && midFixNeedFlipFlag {
		count += 2
	}
	return count
}
