package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	words := strings.Fields(s)
	
	for _,word := range words {
		count,exists  := result[word]
		if !exists {
			result[word] = 0
		}
		result[word] = count + 1
	}
	
	return result
}

func main() {
	wc.Test(WordCount)
}
