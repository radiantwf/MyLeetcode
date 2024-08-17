package main

import (
	"log"
)

func init() {
}

func main() {
	// 输入：edges = [[4,0],[5,4],[1,0],[7,1],[3,7],[6,7],[2,6]]
	log.Println(countGoodNodes([][]int{{4, 0}, {5, 4}, {1, 0}, {7, 1}, {3, 7}, {6, 7}, {2, 6}})) // 4
}

func countGoodNodes(edges [][]int) int {
	nodes := make([][]int, len(edges)+1)
	for _, edge := range edges {
		nodes[edge[0]] = append(nodes[edge[0]], edge[1])
		nodes[edge[1]] = append(nodes[edge[1]], edge[0])
	}
	counter := 0
	var dfs func(int, int) int
	dfs = func(index int, father int) int {
		subtreeChildren := -1
		goodNode := true
		sub := len(nodes[index]) - 1
		for _, childIndex := range nodes[index] {
			if childIndex == father {
				continue
			}

			if subtreeChildren == -1 {
				subtreeChildren = dfs(childIndex, index)
				sub += subtreeChildren
			} else {
				s := dfs(childIndex, index)
				sub += s
				if s != subtreeChildren {
					goodNode = false
				}
			}
		}
		if goodNode {
			counter++
		}
		return sub
	}
	dfs(0, -1)

	return counter
}
