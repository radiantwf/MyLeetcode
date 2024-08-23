package main

import "fmt"

func init() {
}

func main() {
	fmt.Println()
}

func canJump(nums []int) bool {
	reach := 0
	for i, v := range nums {
		if i > reach {
			return false
		}
		reach = max(reach, i+v)
	}
	return true
}
