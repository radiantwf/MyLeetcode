package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minMovesToCaptureTheQueen(1, 6, 3, 3, 5, 6))
}

// 解题思路：
// 皇后和车在同一行或者同一列，只需要一步就可以捕获，中间象不能阻挡
// 皇后和象在同一斜线，只需要一步就可以捕获，中间车不能阻挡
// 其他情况需要2步才能捕获

func minMovesToCaptureTheQueen(a int, b int, c int, d int, e int, f int) int {
	if a == e {
		// 若d在b,f之间
		if (d > b && d < f && a == c) || (d < b && d > f && a == c) {
			return 2
		}
		return 1
	} else if b == f {
		// 若c在a,e之间
		if (c > a && c < e && b == d) || (c < a && c > e && b == d) {
			return 2
		}
		return 1
	} else if math.Abs(float64(c-e)) == math.Abs(float64(d-f)) {
		if math.Abs(float64(c-a)) == math.Abs(float64(d-b)) {
			if c > a && a > e && d > b && b > f {
				return 2
			} else if c < a && a < e && d < b && b < f {
				return 2
			} else if c > a && a > e && d < b && b < f {
				return 2
			} else if c < a && a < e && d > b && b > f {
				return 2
			}
		}
		return 1
	}

	return 2
}
