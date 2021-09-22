package piscine

import (
	//	"fmt"
	//"sort"
	"strings"
)

func MaxWordCountN(text string, n int) map[string]int {
	wordList := strings.Fields(text)
	counts := make(map[string]int)
	for _, word := range wordList {
		_, ok := counts[word]
		if ok {
			counts[word] += 1
		} else {
			counts[word] = 1
		}
	}
	return counts
}
