package main

import "fmt"

func main() {
	fmt.Println(orchestraLayout(3, 0, 2))
	fmt.Println(orchestraLayout(4, 1, 2))
	// fmt.Println(orchestraLayout(5, 2, 2))
}

// 解题思路
// 最外圈 总数量(num-1)*4
// 第2圈 总数量计算： (num-2-1)*4
// 第n圈 总数量计算： (num-2*n-1)*4

// 根据坐标位置，计算所在圈数
// 计算所在圈数的总数量
// 计算所在圈数的起始位置
// 计算所在圈数的结束位置
// 计算总数量
// 总数量%9，如果为余数 0，则返回9，否则返回 余数

func orchestraLayout(num int, xPos int, yPos int) int {
	top, bottom, left, right := xPos+1, num-xPos, yPos+1, num-yPos

	circle := min(left, right, top, bottom)
	counter := 0
	for i := 1; i < circle; i++ {
		counter += (num - (i-1)<<1 - 1) << 2
	}
	if top == circle {
		counter += left - (circle - 1)
	} else if right == circle {
		counter += num - 2*(circle-1) - 1
		counter += top - (circle - 1)
	} else if bottom == circle {
		counter += (num - 2*(circle-1) - 1) * 2
		counter += right - (circle - 1)
	} else if left == circle {
		counter += (num - 2*(circle-1) - 1) * 3
		counter += bottom - (circle - 1)
	}
	ret := counter % 9
	if ret == 0 {
		ret = 9
	}
	return ret
}
