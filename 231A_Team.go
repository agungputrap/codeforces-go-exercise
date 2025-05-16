package main

import "fmt"

const minVotesToImplement = 2

type Vote int

type Problem struct {
	petya Vote
	vasya Vote
	tonya Vote
}

// Problem: 231A - Team
// Source: https://codeforces.com/problemset/problem/231/A
// Solution: Count the number of positive votes for each team.
// If the number of positive votes is greater than or equal to minVotesToImplement,
// then the solution is implementable.
func main() {
	var numProblems int
	fmt.Scan(&numProblems)
	problems := make([]Problem, numProblems)
	for i := 0; i < numProblems; i++ {
		var p, v, t int
		fmt.Scan(&p, &v, &t)
		problems[i] = Problem{Vote(p), Vote(v), Vote(t)}
	}

	result := countImplementableSolutions(problems)
	fmt.Printf("%v\n", result)
}

// countImplementableSolutions counts the number of implementable solutions.
func countImplementableSolutions(problems []Problem) int {
	implementableSolutions := 0

	for _, problem := range problems {
		votes := countPositiveVotes(problem)
		if votes >= minVotesToImplement {
			implementableSolutions++
		}
	}

	return implementableSolutions
}

// countPositiveVotes counts the number of positive votes for each team.
func countPositiveVotes(problem Problem) int {
	return int(problem.petya) + int(problem.vasya) + int(problem.tonya)
}
