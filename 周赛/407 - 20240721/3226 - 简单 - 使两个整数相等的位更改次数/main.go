package main

func main() {
	println(minChanges(13, 4))
}

func minChanges(n int, k int) int {
	if (n | k) != n {
		return -1
	}

	k = n ^ k
	count := 0
	for k > 0 {
		count += k & 1
		k >>= 1
	}
	return count
}
