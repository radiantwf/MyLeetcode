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
// 使用并查集，查询时，把路径上的节点合并，最后路径上的联通块的数量，n-1-联通块数量就是最短距离
// 后续查询时，保留之前联通块数量统计结果，继续合并路径上节点

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	fa := make([]int, n-1)
	for i := range fa {
		fa[i] = i
	}
	// 非递归并查集
	find := func(x int) int {
		rt := x
		for fa[rt] != rt {
			rt = fa[rt]
		}
		for fa[x] != rt {
			fa[x], x = rt, fa[x]
		}
		return rt
	}

	ans := make([]int, len(queries))
	cnt := n - 1
	for qi, q := range queries {
		s, e := q[0], q[1]-1
		fe := find(e)
		for i := find(s); i < e; i = find(i + 1) {
			fa[i] = fe
			cnt--
		}
		ans[qi] = cnt
	}
	return ans
}
