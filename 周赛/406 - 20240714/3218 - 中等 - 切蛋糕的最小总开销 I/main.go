package main

func main() {
	// m = 3, n = 2, horizontalCut = [1,3], verticalCut = [5]
	m := 3
	n := 2
	horizontalCut := []int{1, 3}
	verticalCut := []int{5}
	println(minimumCost(m, n, horizontalCut, verticalCut))
}

// 贪心算法
// 1. 对 horizontalCut 和 verticalCut 进行排序
// 2. 从头开始遍历 horizontalCut 和 verticalCut，每次取最大的一方，计算当前的 cost
// 3. 累加 cost

func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) int {
	quickSort(horizontalCut, 0, len(horizontalCut)-1)
	quickSort(verticalCut, 0, len(verticalCut)-1)
	hIndex := 0
	vIndex := 0
	total := 0
	for hIndex < len(horizontalCut) || vIndex < len(verticalCut) {
		times := 0
		cost := 0
		if hIndex == len(horizontalCut) {
			times = hIndex + 1
			cost = verticalCut[vIndex]
			vIndex++
		} else if vIndex == len(verticalCut) {
			times = vIndex + 1
			cost = horizontalCut[hIndex]
			hIndex++
		} else if verticalCut[vIndex] >= horizontalCut[hIndex] {
			times = hIndex + 1
			cost = verticalCut[vIndex]
			vIndex++
		} else {
			times = vIndex + 1
			cost = horizontalCut[hIndex]
			hIndex++
		}
		total += times * cost
	}

	return total
}

func quickSort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)
		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] > pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
