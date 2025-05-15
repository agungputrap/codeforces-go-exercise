package main

import "fmt"

// Problem: 4A - Watermelon
// Source: https://codeforces.com/problemset/problem/4/A
// Solution: For a watermelon to be divided into two even parts, it must be an even number greater than 2.
// We can check if the number is even and greater than 2.
func main() {
	var watermelon int
	fmt.Scan(&watermelon)
	if watermelon%2 == 0 && watermelon > 2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
