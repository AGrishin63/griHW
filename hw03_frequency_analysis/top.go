package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"fmt"

	s "strings"
)

const top int = 10

func Top10(text string) []string {
	// Place your code here
	text = s.Replace(text, "\n", " ", -1)
	words := s.Split(text, " ")
	dict := make(map[string]int)
	var topVal []int = make([]int, 0, top)
	var topKey []string = make([]string, 0, top)
	var topMin int = 0
	//Заполнить частотный словарь.
	fmt.Println(text)
	for j := 0; j < len(words); j++ {
		word := s.Trim(words[j], " \t")
		if word != "" {
			v := dict[word]
			v++
			dict[word] = v
		} //if
	} //for
	//выбрать топ 10 упорядоченных по убыванию
	for word, v := range dict {
		if v > topMin {
			//занести в тор
			var i int
			for i = 0; i < len(topVal); i++ {
				if v > topVal[i] {
					tmp1 := v
					v = topVal[i]
					topVal[i] = tmp1
					tmp2 := word
					word = topKey[i]
					topKey[i] = tmp2
				} //if
			} //for
			if i < top {
				topVal = append(topVal, v)
				topKey = append(topKey, word)
			} else {
				topMin = v
			} //if
		} //if
	} //for

	return topKey
}
