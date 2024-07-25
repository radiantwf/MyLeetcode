package main

import "fmt"

func main() {
	fmt.Println(losingPlayer(2, 7))
}

// 解题思路：
// 由于要拿取面值总和115的硬币，而硬币的只有10与75两种面值。所以只有一种方案就是拿1个75的硬币与4个10的硬币。
// 所以只需要计算x/1与y/4的最小值即可
// 判断这个值的奇偶性，如果是奇数，则返回Alice，否则返回Bob

func losingPlayer(x int, y int) string {
	v := min(x, y/4)
	if v&1 == 1 {
		return "Alice"
	} else {
		return "Bob"
	}
}
