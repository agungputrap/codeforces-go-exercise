package main

import "fmt"

// Problem: 1A - Theatre Square
// Source: https://codeforces.com/problemset/problem/1/A
// Solution: count how many tiles are needed in each direction with ceiling value
func main() {
	var n, m, a int
	fmt.Scan(&n, &m, &a)
	fmt.Println(countTiles(n, m, a))
}

// countTiles calculate how many tiles are needed in each direction
func countTiles(n, m, a int) int {
	tilesN := (n + a - 1) / a
	tilesM := (m + a - 1) / a
	return tilesN * tilesM
}
