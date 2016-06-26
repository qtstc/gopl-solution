package main

import "fmt"

func main() {
	fmt.Println(isAnagram("abac", "cbca"))
}

func isAnagram(s1, s2 string) bool {
	return getCharCount(s1) == getCharCount(s2)
}

func getCharCount(s string) [26]byte {
	var count [26]byte
	for _, char := range s {
		count[char-'a']++
	}
	return count
}
