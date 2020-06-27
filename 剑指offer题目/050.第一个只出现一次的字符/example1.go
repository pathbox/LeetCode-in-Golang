package main

import "fmt"

func firstUniqChar(s string) byte {
	m := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			m[s[i]] = false
		} else {
			m[s[i]] = true
		}
	}
	for i := 0; i < len(s); i++ {
		if m[s[i]] {
			return s[i]
		}
	}
	return byte(' ')
}

func main() {
	s := "aadadaad"
	fmt.Println(firstUniqChar(s))
}
