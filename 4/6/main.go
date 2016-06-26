package main

import "fmt"
import "unicode"

func main() {
	fmt.Println(string(squashSpace([]byte("　h 　"))))
}

func squashSpace(s []byte) []byte {
	squashed, count := s[:0], 0
	for _, c := range string(s) {
		if unicode.IsSpace(c) {
			squashed = append(squashed, ' ')
			count++
		} else {
			for _, c := range []byte(string(c)) {
				squashed = append(squashed, c)
				count++
			}
		}
	}
	return squashed[:count]
}
