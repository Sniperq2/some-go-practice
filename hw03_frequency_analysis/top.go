package hw03frequencyanalysis

import (
	"fmt"
	"sort"
	"strings"
)

func Top10(line string) []string {
	wordCounter := make(map[string]int)
	splitter := strings.Fields(line)

	for _, word := range splitter {
		lightWord := strings.TrimSpace(word)

		_, found := wordCounter[lightWord]

		if found && len(lightWord) > 0 {
			wordCounter[lightWord] += 1
		} else {
			wordCounter[lightWord] = 1
		}
	}

	keys := make([]string, 0, len(wordCounter))

	for k := range wordCounter {
		keys = append(keys, k)
	}

	fmt.Println(keys)

	sort.SliceStable(keys, func(i, j int) bool {
		return wordCounter[keys[i]] > wordCounter[keys[j]]
	})

	var h = make([]string, 0, 10)

	for idx, k := range keys {
		if idx < 10 {
			h = append(h, k)
		}
	}

	return h
}
