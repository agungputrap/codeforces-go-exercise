package main

import "fmt"

const maxWordLength = 10

// Problem: 71A - Way Too Long Words
// Source: https://codeforces.com/problemset/problem/71/A
// Solution: If a word longer than maxWordLength, it's abbreviated to:
// first character + number of characters between first and last character + last character
func main() {
	var wordCount int
	fmt.Scan(&wordCount)
	words := readWords(wordCount)
	for _, word := range words {
		fmt.Println(processWord(word))
	}
}

// readWords reads a specified number of words from standard input
func readWords(count int) []string {
	var words []string
	for i := 0; i < count; i++ {
		var word string
		fmt.Scan(&word)
		words = append(words, word)
	}
	return words
}

// processWord handles the word abbreviation logic
func processWord(word string) string {
	if len(word) <= maxWordLength {
		return word
	}
	firstChar := word[0:1]
	lastChar := word[len(word)-1:]
	middleLength := len(word) - 2
	return fmt.Sprintf("%s%v%s", firstChar, middleLength, lastChar)
}
