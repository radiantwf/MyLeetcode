package main

import "fmt"

func main() {
	fmt.Println()
}

// 解题思路
// 只有在所有个位数与所有两位数的和相等时，A才会输
func canAliceWin(nums []int) bool {
	sum1, sum2 := 0, 0
	for _, num := range nums {
		if num < 10 {
			sum1 += num
		} else {
			sum2 += num
		}
	}
	return sum1 != sum2
}
