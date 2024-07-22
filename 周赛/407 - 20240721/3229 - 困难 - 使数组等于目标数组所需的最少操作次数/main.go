package main

func main() {
	nums := []int{0, 0, 0, 0}
	target := []int{-1, -3, -3, -1}

	println(minimumOperations(nums, target))
}

// 解题思路
// 1、首先计算target与nums每个元素的差值，构造新的数组
// 2、对新的数组计算差分数组
// 3、计算"差分数组变为0,0,0......0"，所需的最小操作次数

func minimumOperations(nums []int, target []int) int64 {
	d_pre := target[0] - nums[0]
	times := int64(max(d_pre, 0))

	for i := 1; i < len(nums); i++ {
		d_i := target[i] - nums[i]
		d := d_i - d_pre
		times += int64(max(d, 0))
		d_pre = d_i
	}
	times += int64(max(-d_pre, 0))
	return times
}
