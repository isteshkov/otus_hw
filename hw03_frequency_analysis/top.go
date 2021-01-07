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
		if wordsSort[j].Count == wordsSort[i].Count {
			return true
		}
		return wordsSort[j].Count < wordsSort[i].Count
	})

	top := make([]string, 0, 10)
	for i := 0; i < 10; i++ {
		top = append(top, wordsSort[i].Value)
	}

	return top
}
