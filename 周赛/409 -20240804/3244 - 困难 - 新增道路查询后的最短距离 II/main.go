package main

import (
	"fmt"
	"math/bits"
)

func init() {
}

func main() {
	fmt.Println(shortestDistanceAfterQueries(12, [][]int{{5, 9}, {7, 9}, {9, 11}}))
}

// 解题思路
// 所有查询中不会存在两个查询都满足 queries[i][0] < queries[j][0] < queries[i][1] < queries[j][1]
// 得出道路不会有交叉
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
	cnt := n - 1 // 并查集连通块个数
	for qi, q := range queries {
		l, r := q[0], q[1]-1
		fr := find(r)
		for i := find(l); i < r; i = find(i + 1) {
			fa[i] = fr
			cnt--
		}
		ans[qi] = cnt
	}
	return ans
}

func shortestDistanceAfterQueries2(n int, queries [][]int) []int {
	// 3 <= n <= 500
	nodeBits := make([]uint64, n/64+1)
	// 每一位，表示一个节点，把所有节点置1
	for i := range nodeBits {
		if (i+1)*64 <= n {
			nodeBits[i] = 0xffffffffffffffff
		} else if i*64 < n {
			t := n - i*64
			for j := 0; j < t; j++ {
				// 从高位向低位，置1
				nodeBits[i] |= 1 << uint(63-j)
			}
			break
		} else {
			break
		}
	}

	addRoad := func(start, end int) {
		startByteListIndex := (start + 1) / 64
		startBitIndex := (start + 1) % 64
		endByteListIndex := (end - 1) / 64
		endBitIndex := (end - 1) % 64
		// 从 start 到 end 的所有节点，置0
		for i := startByteListIndex; i <= endByteListIndex; i++ {
			if i == startByteListIndex {
				end := 63
				if i == endByteListIndex {
					end = endBitIndex
				}
				mask1 := uint64(0xffffffffffffffff) << uint(64-startBitIndex)
				mask2 := uint64(0xffffffffffffffff) >> uint(end+1)
				mask := mask1 | mask2
				nodeBits[startByteListIndex] &= mask
			} else if i == endByteListIndex {
				start := 0
				mask1 := uint64(0xffffffffffffffff) << uint(64-start)
				mask2 := uint64(0xffffffffffffffff) >> uint(endBitIndex+1)
				mask := mask1 | mask2
				nodeBits[endByteListIndex] &= mask
			} else {
				nodeBits[i] = 0
			}
		}
	}
	getBitOneCount := func() int {
		count := 0
		for _, b := range nodeBits {
			count += bits.OnesCount64(b)
		}
		return count
	}
	results := make([]int, len(queries))
	for i, road := range queries {
		addRoad(road[0], road[1])
		results[i] = getBitOneCount() - 1
	}

	return results
}
