package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("12345678"))
}

func comma(s string) string {
	var buff bytes.Buffer
	for i, j := len(s)-1, 0; i >= 0; i-- {
		if j == 3 {
			buff.WriteByte(',')
			j = 0
		}
		j++
		buff.WriteByte(s[i])
	}
	return reverse(buff.String())
}

func reverse(s string) string {
	bytes := []byte(s)
	for i, j := 0, len(bytes)-1; i < j; {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	return string(bytes)
}
