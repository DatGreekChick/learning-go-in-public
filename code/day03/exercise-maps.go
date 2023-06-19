package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	counts := map[string]int{}

	for i := range words {
		word := words[i]
		counts[word] += 1
	}

	return counts
}

func main() {
	wc.Test(WordCount)
}
