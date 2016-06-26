package main

import "fmt"

func main() {
	a := [3]int{3, 2, 1}
	reverse(&a)
	fmt.Println(a)
}

func reverse(a *[3]int) {
	for i, j := 0, len(*a)-1; i < j; i, j = i+1, j-1 {
		(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
	}
}
