package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("-1312314231.008"))
}

func comma(s string) string {
	if s[0] == '-' {
		return "-" + commaForFloat(s[1:])
	}
	return commaForFloat(s)
}

func commaForFloat(s string) string {
	dotIndex := strings.LastIndex(s, ".")
	if dotIndex >= 0 {
		return commaForInt(s[:dotIndex]) + s[dotIndex:]
	}
	return commaForInt(s)
}

func commaForInt(s string) string {
	if len(s) <= 3 {
		return s
	}
	return commaForInt(s[:len(s)-3]) + "," + s[len(s)-3:]
}
