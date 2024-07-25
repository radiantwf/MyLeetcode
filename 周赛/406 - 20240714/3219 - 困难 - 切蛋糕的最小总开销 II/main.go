package main

func main() {
	m := -1
	n := -1
	horizontalCut := []int{333, 1, 1, 333}
	verticalCut := []int{1, 2}
	println(minimumCost(m, n, horizontalCut, verticalCut))
}

// 解题思路
// 由于：
// 1 <= m, n <= 10^5
// 1 <= horizontalCut[i], verticalCut[i] <= 10^3
// 所以 horizontalCut 和 verticalCut 可能出现大量重复值，所以先对这两个数组进行计数器排序，并保存数量
// 然后按照贪心算法的思路，每次取最大的一方，计算当前的 cost，累加 cost

func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) int64 {
	h := countSortAndMerge(horizontalCut)
	v := countSortAndMerge(verticalCut)
	hIndex := 0
	vIndex := 0
	hCutTimes := 1
	vCutTimes := 1
	total := int64(0)
	for hIndex < len(h) || vIndex < len(v) {
		times := 0
		cost := 0
		if hIndex == len(h) {
			times = vCutTimes * v[vIndex].Times
			cost = v[vIndex].Value
			hCutTimes += v[vIndex].Times
			vIndex++
		} else if vIndex == len(v) {
			times = hCutTimes * h[hIndex].Times
			cost = h[hIndex].Value
			vCutTimes += h[hIndex].Times
			hIndex++
		} else if v[vIndex].Value >= h[hIndex].Value {
			times = vCutTimes * v[vIndex].Times
			cost = v[vIndex].Value
			hCutTimes += v[vIndex].Times
			vIndex++
		} else {
			times = hCutTimes * h[hIndex].Times
			cost = h[hIndex].Value
			vCutTimes += h[hIndex].Times
			hIndex++
		}
		total += int64(times * cost)
	}

	return total
}

type ValueTimes struct {
	Value int
	Times int
}

func countSortAndMerge(arr []int) []ValueTimes {
	if len(arr) == 0 {
		return nil
	}

	// 找到数组中的最大值和最小值
	maxVal := arr[0]
	minVal := arr[0]
	for _, v := range arr {
		if v > maxVal {
			maxVal = v
		}
		if v < minVal {
			minVal = v
		}
	}

	// 创建计数数组
	count := make([]int, maxVal-minVal+1)

	// 填充计数数组
	for _, v := range arr {
		count[v-minVal]++
	}

	// 创建结果数组
	var result []ValueTimes

	// 从大到小遍历计数数组并填充结果数组
	for i := len(count) - 1; i >= 0; i-- {
		if count[i] > 0 {
			result = append(result, ValueTimes{Value: i + minVal, Times: count[i]})
		}
	}

	return result
}

// import "sort"
// func countSortAndMerge(arr []int) []ValueTimes {
// 	if len(arr) == 0 {
// 		return nil
// 	}

// 	// 使用 map 记录每个数值的出现次数
// 	count := make(map[int]int)
// 	for _, v := range arr {
// 		count[v]++
// 	}

// 	// 创建结果数组
// 	var result []ValueTimes

// 	// 将 map 中的键值对转换为结果数组，并按值从大到小排序
// 	for value, times := range count {
// 		result = append(result, ValueTimes{Value: value, Times: times})
// 	}

// 	// 按值从大到小排序
// 	sort.Slice(result, func(i, j int) bool {
// 		return result[i].Value > result[j].Value
// 	})

// 	return result
// }
