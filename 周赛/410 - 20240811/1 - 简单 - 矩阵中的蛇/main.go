package main

import "fmt"

func init() {
}

func main() {
	// 输入：n = 3, commands = ["DOWN","RIGHT","UP"]
	fmt.Println(finalPositionOfSnake(3, []string{"DOWN", "RIGHT", "UP"}))
}

func finalPositionOfSnake(n int, commands []string) int {
	i, j := 0, 0
	for _, command := range commands {
		switch command[0] {
		case 'U':
			i--
		case 'D':
			i++
		case 'L':
			j--
		default:
			j++
			break
		}
	}

	return i*n + j
}
