package main

import (
	"fmt"
	"math"
)

func init() {
}

func main() {
	fmt.Println(shortestDistanceAfterQueries(5, [][]int{{2, 4}, {0, 2}, {0, 4}}))
}

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	graphNodes := make([]*GraphNode, n)
	graphNodes[0] = &GraphNode{Val: 0}
	for i := 1; i < n; i++ {
		graphNodes[i] = &GraphNode{Val: i}
		graphNodes[i-1].Next = append(graphNodes[i-1].Next, graphNodes[i])
	}
	results := make([]int, len(queries))
	dfs := func() int {
		dist := make([]int, n)
		for i := range dist {
			dist[i] = math.MaxInt32
		}
		dist[0] = 0

		queue := []int{0}
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			for _, neighbor := range graphNodes[node].Next {
				if dist[neighbor.Val] == math.MaxInt32 {
					dist[neighbor.Val] = dist[node] + 1
					queue = append(queue, neighbor.Val)
					if neighbor.Val == n-1 {
						return dist[neighbor.Val]
					}
				}
			}
		}
		return -1

	}
	for i, road := range queries {
		graphNodes[road[0]].Next = append(graphNodes[road[0]].Next, graphNodes[road[1]])
		results[i] = dfs()
	}
	return results
}

type GraphNode struct {
	Val  int
	Next []*GraphNode
}
