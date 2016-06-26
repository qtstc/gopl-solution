package main

import "fmt"

func main() {
	s := "śá1哈21æ2âà3和â123"
	bytes := []byte(s)
	fmt.Println(string(reverse(bytes)))
}

func reverse(bytes []byte) []byte {
	s := []rune(string(bytes))
	for si, sj, bi, bj, li, lj := 0, len(s)-1, 0, len(bytes)-1, 0, 0; ; {
		if si > sj {
			return bytes
		}
		if li == 0 && lj == 0 {
			li = len(string(s[si]))
			lj = len(string(s[sj]))
			si++
			sj--
		} else if li == 0 {
			li = len(string(s[si]))
			si++
		} else {
			lj = len(string(s[sj]))
			sj--
		}
		bi, bj, li, lj, bytes = reverseRecur(bytes, bi, bj, li, lj)
	}
}

func reverseRecur(bytes []byte, i int, j int, li int, lj int) (int, int, int, int, []byte) {
	if li == lj {
		return i + li, j - lj, 0, 0, reverseSameLen(bytes, i, j-li+1, li)
	}
	if li < lj {
		_, _, _, _, bytes = reverseRecur(bytes, j-lj+1, j, li, lj-li)
		return i + li, j - li, 0, lj - li, reverseSameLen(bytes, i, j-li+1, li)
	}
	_, _, _, _, bytes = reverseRecur(bytes, i, i+li-1, li-lj, lj)
	return i + lj, j - lj, li - lj, 0, reverseSameLen(bytes, i, j-lj+1, lj)
}

func reverseSameLen(bytes []byte, i int, j int, l int) []byte {
	for n := 0; n < l; i, j, n = i+1, j+1, n+1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return bytes
}
