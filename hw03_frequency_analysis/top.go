package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"sort"
	"strings"
)

type Word struct {
	Value string
	Count int
}

func Top10(s string) []string {
	if len(s) == 0 {
		return []string{}
	}

	words := strings.Fields(s)
	wordsCounts := make(map[string]int)

	for i := range words {
		wordsCounts[words[i]]++
	}

	wordsSort := make([]Word, 0, len(wordsCounts))
	for w, c := range wordsCounts {
		wordsSort = append(wordsSort, Word{
			Value: w,
			Count: c,
		})
	}

	sort.Slice(wordsSort, func(i, j int) bool {
		if wordsSort[i].Count == wordsSort[j].Count {
			return wordsSort[i].Value > wordsSort[j].Value
		}
		return wordsSort[i].Count > wordsSort[j].Count
	})

	top := make([]string, 0, 10)
	for i := 0; i < len(wordsSort); i++ {
		top = append(top, wordsSort[i].Value)
		if len(top) == 10 {
			break
		}
	}

	return top
}
