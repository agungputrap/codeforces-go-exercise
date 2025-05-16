package main

import "fmt"

const (
	incrementPrefix  = "++X"
	incrementPostfix = "X++"
	decrementPrefix  = "--X"
	decrementPostfix = "X--"
)

type Operation string

// Problem: 282A - Bit Plus Plus
// Source: https://codeforces.com/problemset/problem/282/A
// Solution: Iterate through the operations and increment/decrement the result based on the operation.
func main() {
	var numOperations int
	fmt.Scan(&numOperations)
	operations := make([]Operation, numOperations)
	for i := 0; i < numOperations; i++ {
		var operation Operation
		fmt.Scan(&operation)
		operations[i] = operation
	}
	result := calculateFinalValue(operations)
	fmt.Printf("%v\n", result)
}

// calculateFinalValue calculates the final value based on the operations.
func calculateFinalValue(operations []Operation) int {
	var result int
	for _, operation := range operations {
		switch operation {
		case incrementPrefix:
			result++
		case incrementPostfix:
			result++
		case decrementPrefix:
			result--
		case decrementPostfix:
			result--
		}
	}
	return result
}
