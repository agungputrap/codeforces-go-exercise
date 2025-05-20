package main

import "fmt"

const (
	width = 2
)

// Problem: 50A - Domino Pilling
// Source: https://codeforces.com/problemset/problem/50/A
// Solution: Calculate maximum number of 2×1 dominoes that can be placed on an M×N board.
// When M is even, we can fill all rows with horizontal dominoes.
// When M is odd, we fill M-1 rows with horizontal dominoes and the last row with vertical dominoes.

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	fmt.Println(Pilling(m, n))
}

// Pilling calculates the maximum number of 2×1 dominoes that can be placed on an M×N board.
func Pilling(m, n int) int {
	if m%2 == 0 {
		return (m / width) * n
	}
	return ((m / width) * n) + (n / width)
}
