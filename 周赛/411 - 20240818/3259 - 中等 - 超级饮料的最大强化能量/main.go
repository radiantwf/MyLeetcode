package main

import (
	"fmt"
)

func init() {
}

func main() {
	fmt.Println(maxEnergyBoost(
		[]int{30, 45, 21, 96, 73, 48, 67},
		[]int{84, 47, 52, 50, 10, 65, 14}))
}

func maxEnergyBoost(energyDrinkA []int, energyDrinkB []int) int64 {
	arrs := [2][]int{energyDrinkA, energyDrinkB}
	n := len(energyDrinkA)
	values := [2][2]int64{}
	index := 0
	for i := 0; i < n; i++ {
		index = i & 1
		values[index][0], values[index][1] = max(values[^index&1][0], values[index][1])+int64(arrs[0][i]), max(values[^index&1][1], values[index][0])+int64(arrs[1][i])
	}
	return max(values[index][0], values[index][1])
}
