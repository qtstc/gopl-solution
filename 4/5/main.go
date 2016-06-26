package main

import "fmt"

func main() {
	fmt.Println(dedup([]string{"hehe", "haha", "haha", "hehe"}))
}

func dedup(a []string) []string {
	p := 1
	for i := 1; i < len(a); i++ {
		if a[i-1] != a[i] {
			a[p] = a[i]
			p++
		}
	}
	return a[:p]
}
