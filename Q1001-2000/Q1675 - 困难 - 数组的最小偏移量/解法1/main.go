package main

import (
	"container/heap"
	"log"
)

func main() {
	nums := []int{3, 5}

	v := minimumDeviation(nums)
	log.Println(v)
}

// 解题思路
// 1. 把所有奇数*2，所有数值会扩展到最大值（奇数*2后会变为偶数，偶数只可以进行/2操作）
// 2. 操作1遍历数组时，同时找到最小值
// 3. 把数组转换为最大堆
// 4. 从堆顶开始，如果堆顶是偶数，就/2，检查最小值，然后重新调整堆
// 5. 比较堆顶和最小值的差值，如果比之前的差值小，就更新差值
// 6. 如果堆顶是奇数，就停止，返回最小值和堆顶的差值
// 7. 把*2,/2,%2的操作都转换为位运算，效率更高

// import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minimumDeviation(nums []int) int {
	min := int(^uint(0) >> 1)
	for i := 0; i < len(nums); i++ {
		if nums[i]&1 == 1 {
			nums[i] <<= 1
		}
		if nums[i] < min {
			min = nums[i]
		}
	}
	h := IntHeap(nums)
	heap.Init(&h)

	d := h[0] - min
	for h[0]&1 == 0 {
		h[0] >>= 1
		if h[0] < min {
			min = h[0]
		}
		heap.Fix(&h, 0)
		tmp := h[0] - min
		if tmp <= d {
			d = tmp
		}
	}
	return d
}
