package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
    count := make(map[string]int)

	for w := range words {
		count[words[w]] += 1
	}
	
	return count
}

func main() {
	wc.Test(WordCount)
}
