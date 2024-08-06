package main

import (
	"fmt"
)

func init() {
}

func main() {
	fmt.Println(shortestDistanceAfterQueries(12, [][]int{{5, 9}, {7, 9}, {9, 11}}))
}

// 解题思路
// 所有查询中不会存在两个查询都满足 queries[i][0] < queries[j][0] < queries[i][1] < queries[j][1]
// 得出道路不会有交叉
// 定义一个维度为n的数组，每个元素表示当前节点节约（跨越）的距离，同时定义一个变量，表示总共节约的距离
// 每次查询，设置节约的距离，与总节约距离的变量（需要遍历节约距离数据是否有重叠部分）
// 最短距离等于n-1-总节约距离

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	saveDistanceList := make([]int, n)
	totalSave := 0
	addRoad := func(s int, e int) {
		save := e - s - 1
		if saveDistanceList[s] > save {
			return
		}
		if saveDistanceList[s] < 0 {
			return
		}
		totalSave = totalSave - saveDistanceList[s] + save
		saveDistanceList[s] = save
		for i := s + 1; i < e; i++ {
			if saveDistanceList[i] > 0 {
				totalSave -= saveDistanceList[i]
			}
			saveDistanceList[i] = -1
		}
	}
	results := make([]int, len(queries))
	for i, road := range queries {
		addRoad(road[0], road[1])
		results[i] = n - 1 - totalSave
	}
	return results
}
