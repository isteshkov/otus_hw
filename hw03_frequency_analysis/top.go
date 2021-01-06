package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"regexp"
	"sort"
	"strings"
)

var replaceRx = regexp.MustCompile(`[\t\s]+`)

type Word struct {
	Value string
	Count int
}

func Top10(s string) []string {
	// Place your code here

	if len(s) == 0 {
		return []string{}
	}

	wordsCounts := make(map[string]int)

	s = strings.TrimSpace(replaceRx.ReplaceAllString(s, " "))
	lines := strings.Split(s, "\n")

	words := make([]string, 0)
	for i := range lines {
		words = append(words, strings.Split(lines[i], " ")...)
	}
	for i := range words {
		wordsCounts[words[i]]++
	}

	wordsSort := make([]Word, 0, 10)
	for w, c := range wordsCounts {
		wordsSort = append(wordsSort, Word{
			Value: w,
			Count: c,
		})
	}

	sort.Slice(wordsSort, func(i, j int) bool {
		return wordsSort[j].Count < wordsSort[i].Count
	})

	top := make([]string, 0, 10)
	for i := range wordsSort {
		top = append(top, wordsSort[i].Value)
		if i == 9 {
			break
		}
	}

	return top
}
