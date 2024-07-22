package main

import "fmt"

func main() {
	s := "45320"
	fmt.Println(getSmallestString(s))
}

func getSmallestString(s string) string {
	b := []byte(s)
	prev := b[0]
	for i := 1; i < len(b); i++ {
		current := b[i]
		if ((prev^current)&1 | 0) == 0 {
			if current < prev {
				b[i-1], b[i] = b[i], b[i-1]
				break
			}
		}
		prev = current
	}

	return string(b)
}
