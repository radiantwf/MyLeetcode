package main

import (
	"fmt"
	"log"
)

func numberOfAlternatingGroups(colors []int, queries [][]int) []int {
	adjacent := getAlternatingFlags(colors)
	adjacentSegmentTree := make([]int, len(adjacent)*4)
	buildSegmentTree(adjacentSegmentTree, adjacent, 0, len(adjacent)-1, 0)
	log.Printf("%v", adjacentSegmentTree)

	for _, query := range queries {
		if query[0] == 2 {
			if colors[query[1]] == query[2] {
				continue
			}
			colors[query[1]] = query[2]
			index0, index1 := query[0]-1, query[0]
			if index0 < 0 {
				index0 = len(colors) - 1
			}
			updateSegmentTree(adjacentSegmentTree, 0, len(adjacent)-1, 0, index0)
			updateSegmentTree(adjacentSegmentTree, 0, len(adjacent)-1, 0, index1)
			log.Printf("%v", adjacentSegmentTree)
		} else if query[0] == 1 {
			// checkValue := query[1]
		}

	}

	results := []int{}
	return results
}

func main() {
	colors := []int{0, 0, 1, 0, 1, 1}
	queries := [][]int{{1, 3}, {2, 3, 0}, {1, 5}}
	fmt.Println(numberOfAlternatingGroups(colors, queries)) // 预期输出: [2, 0]
}

// 构造线段树
func buildSegmentTree(segmentTree []int, colors []int, left int, right int, index int) {
	if left == right {
		segmentTree[index] = colors[left]
		return
	}
	mid := left + (right-left)/2
	buildSegmentTree(segmentTree, colors, left, mid, index*2+1)
	buildSegmentTree(segmentTree, colors, mid+1, right, index*2+2)
	segmentTree[index] = segmentTree[index*2+1] + segmentTree[index*2+2]
}

// 更新adjacent数组index-1与index的值，同时更新线段树
func updateSegmentTree(segmentTree []int, left int, right int, index int, updateIndex int) {
	if left == right {
		segmentTree[index] = ^segmentTree[index] & 1
		return
	}
	mid := left + (right-left)/2
	if updateIndex <= mid {
		updateSegmentTree(segmentTree, left, mid, index*2+1, updateIndex)
	} else {
		updateSegmentTree(segmentTree, mid+1, right, index*2+2, updateIndex)
	}
	segmentTree[index] = segmentTree[index*2+1] + segmentTree[index*2+2]
}

func getAlternatingFlags(colors []int) []int {
	result := make([]int, len(colors))
	for i := 0; i < len(colors)-1; i++ {
		result[i] = colors[i] ^ colors[i+1]
	}
	result[len(colors)-1] = colors[len(colors)-1] ^ colors[0]
	return result
}
