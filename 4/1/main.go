package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	fmt.Println(countDiffBits(sha256.Sum256([]byte("x")), sha256.Sum256([]byte("X"))))
}

func countDiffBits(b1, b2 [32]byte) int {
	n := 0
	for i := 0; i < 32; i++ {
		n += popCount(b1[i] ^ b2[i])
	}
	return n
}

func popCount(b byte) int {
	n := 0
	for b != 0 {
		b &= (b - 1)
		n++
	}
	return n
}
