package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type Pair struct {
	word   string
	repeat int
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Top10(str string) []string {
	if len(str) == 0 {
		return nil
	}
	words := strings.Fields(str)
	repeats := make(map[string]int)
	for _, word := range words {
		repeats[word]++
	}
	pairs := make([]Pair, 0, len(repeats))
	for word, repeat := range repeats {
		pairs = append(pairs, Pair{word, repeat})
	}
	sort.SliceStable(pairs, func(i, j int) bool {
		if pairs[i].repeat == pairs[j].repeat {
			return pairs[i].word < pairs[j].word
		}
		return pairs[i].repeat > pairs[j].repeat
	})
	res := make([]string, len(pairs))
	for i := 0; i < len(pairs); i++ {
		res[i] = pairs[i].word
	}
	return res[:min(10, len(res))]
}
