package main

import (
	"fmt"
)

func main() {
	fmt.Println(minChanges([]int{0, 1, 2, 3, 3, 6, 5, 4}, 6))
}

// 解题思路
// 2个数字都修改，则2个数字的差值的绝对值范围为0~k
// 修改2个数字中则1个数字，对应2个数字的差值的绝对值范围为：
// 未修改数字 小于k，对应差值为 0~max(k-x1, x1, k-x2, x2)
// 未修改数字，差值为abs(x-y)

// 定义2个维度为k+1的计数器数组，分别统计：
// 1、对应序列相减后的差值的绝对值的次数
// 2、只修改一个数字，差值绝对值得取值范围可以为[0,x]的数组项（数组0-half_len检索）

func minChanges(nums []int, k int) int {
	differenceCounterList := make([]int, k+1)
	ChangeOnceDifferenceMaxList := make([]int, k+1)
	numslen := len(nums)
	numsHalfLen := numslen / 2
	for i := 0; i < numsHalfLen; i++ {
		if nums[i] < nums[numslen-1-i] {
			differenceCounterList[nums[numslen-1-i]-nums[i]]++
			ChangeOnceDifferenceMaxList[max(k-nums[i], nums[numslen-1-i])]++
		} else {
			differenceCounterList[nums[i]-nums[numslen-1-i]]++
			ChangeOnceDifferenceMaxList[max(nums[i], k-nums[numslen-1-i])]++
		}
	}

	minTimes := numslen
	changeOnceCounter := 0
	for x := 0; x <= k; x++ {
		minTimes = min(numsHalfLen-differenceCounterList[x]+changeOnceCounter, minTimes)
		changeOnceCounter += ChangeOnceDifferenceMaxList[x]
	}
	return minTimes
}
