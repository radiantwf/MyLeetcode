package main

import "fmt"

func init() {
}

func main() {
	fmt.Println()
}

func canJump(nums []int) bool {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == 0 {
			for j := i - 1; j >= 0; j-- {
				if nums[j] > i-j {
					break
				}
				if j == 0 {
					return false
				}
			}
		}
	}
	return true
}
