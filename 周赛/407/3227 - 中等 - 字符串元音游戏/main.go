package main

func main() {
	s := "leetcoder"
	println(doesAliceWin(s))
}

// 题目分析：
// 1、有奇数个元音字母A胜
// 2、有非0偶数个元音字母，由于A先手，A必胜
// 3、有0个元音字母，B胜

func doesAliceWin(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			return true
		}
	}
	return false
}
