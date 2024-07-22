package main

func main() {
	s := "100110100"
	println(maxOperations(s))
}

func maxOperations(s string) int {
	times := 0
	counter := 0
	lastZero := true
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			counter++
			lastZero = false
		} else {
			if !lastZero {
				times += counter
				lastZero = true
			}
		}
	}
	return times
}
