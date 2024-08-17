package main

import "fmt"

func init() {
}

func main() {
	// [5,5,5,5]
	fmt.Println(countOfPairs([]int{5, 5, 5, 5}))
}

func countOfPairs(nums []int) int {
	n := len(nums)
	counter := 0
	var checkFunc func(int, int, int)
	checkFunc = func(index, last1, last2 int) {
		for v1 := 0; v1 <= nums[index]; v1++ {
			v2 := nums[index] - v1
			if index > 0 {
				if v1 < last1 {
					continue
				} else if v2 > last2 {
					continue
				}
			}
			if index == n-1 {
				counter++
				if counter == 1000000007 {
					counter = 0
				}
			} else {
				checkFunc(index+1, v1, v2)
			}
		}
	}
	checkFunc(0, -1, -1)
	return counter
}
