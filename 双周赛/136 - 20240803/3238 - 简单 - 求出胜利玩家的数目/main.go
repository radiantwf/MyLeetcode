package main

import "fmt"

func init() {
}

func main() {
	// n = 5, pick = [[1,1],[2,4],[2,4],[2,4]]
	fmt.Println(winningPlayerCount(5, [][]int{{1, 1}, {2, 4}, {2, 4}, {2, 4}})) // 1
}

func winningPlayerCount(n int, pick [][]int) int {
	counter := make(map[int]int, 100)
	for i := 0; i < len(pick); i++ {
		counter[pick[i][1]*10+pick[i][0]]++
	}
	players := make([]int, n)
	for k, v := range counter {
		player := k % 10
		if v >= player+1 {
			players[player]++
		}
	}
	winners := 0
	for i := 0; i < len(players); i++ {
		if players[i] > 0 {
			winners++
		}
	}
	return winners
}
