package main

import "fmt"

func init() {
}

func main() {
	fmt.Println(resultsArray([]int{1, 3, 4}, 2))
}

/*
 * @lc app=leetcode.cn id=100383 lang=golang
 *
 * [100383] 长度为 K 的子数组的能量值 I
 */

// @lc code=start
func resultsArray(nums []int, k int) []int {
	result := make([]int, len(nums)-k+1)
	counter := 0
	last := -1
	for i, num := range nums {
		if num == last+1 {
			counter++
		} else {
			counter = 1
		}
		last = num
		if i+1-k >= 0 {
			if counter >= k {
				result[i-k+1] = last
			} else {
				result[i-k+1] = -1
			}
		}
	}
	return result
}

// @lc code=end
