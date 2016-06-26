package main

import "fmt"

func main() {
	fmt.Println(rotate([]int{10, 20, 30, 40, 50}, 5))
}

func rotate(a []int, n int) []int {
	for i := 0; i < gcd(len(a), n); i++ {
		for tmp, k := a[i], (i+n)%len(a); ; {
			a[k], tmp = tmp, a[k]
			if i == k {
				break
			}
			k = (k + n) % len(a)
		}
	}
	return a
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
