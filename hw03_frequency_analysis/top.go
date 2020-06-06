package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"sort"
	"strings"
)

const top int = 10

func Top10(text string) []string {
	// Place your code here
	text = strings.Replace(text, "\n", " ", -1)
	words := strings.Split(text, " ")
	dict := make(map[string]int)
	topKey := make([]string, 0, top)

	//Заполнить частотный словарь.
	for j := 0; j < len(words); j++ {
		word := strings.Trim(words[j], " \t")
		if word != "" {
			dict[word]++
		}
	}

	//создать слайс структур
	type item struct {
		wrd string
		frq int
	}
	tmp := make([]item, 0, len(dict))

	//перелить словарь в слайс структур
	for word, v := range dict {
		t := new(item)
		t.wrd = word
		t.frq = v
		tmp = append(tmp, *t)
	}

	// сортировать и взять первые 10
	sort.Slice(tmp, func(i, j int) bool { return tmp[i].frq > tmp[j].frq })
	for m := 0; m < top && m < len(tmp); m++ {
		topKey = append(topKey, tmp[m].wrd)
	}
	return topKey
}
