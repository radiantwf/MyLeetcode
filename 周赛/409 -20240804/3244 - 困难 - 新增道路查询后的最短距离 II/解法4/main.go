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
// 原理与解法3一样，只是使用数组代替并查集
// 定义一个后续节点数组，记录每个节点的下一个节点。定义一个计数器，为初始状态起点到终点的距离（n-1）
// 添加道路时，从道路起点开始，一直向后遍历到道路终点
// 遍历时：后续节点数组的值小于道路终点时，更新后续节点数组的值，同时计数器减一，大于等于时直接跳出循环即可
// 最终返回计数器的值即可

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	next := make([]int, n-1)
	for i := range next {
		next[i] = i + 1
	}
	ans := make([]int, len(queries))
	cnt := n - 1
	for qi, q := range queries {
		for i := q[0]; i < q[1]; {
			if q[1] > next[i] {
				next[i], i = q[1], next[i]
				cnt--
			} else {
				break
			}
		}
		ans[qi] = cnt
	}
	return ans
}
