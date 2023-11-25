package main

import (
	"fmt"
	"sort"
	"strings"
)

func CountWords(text string) map[string]int {
	// Implement Me :)
	mp := make(map[string]int)
	if len(text) > 0 {
		strSlice := strings.Fields(text)
		for _, v := range strSlice {
			//fmt.Println(v)
			if _, exists := mp[v]; exists {
				mp[v] += 1
			} else {
				mp[v] = 1
			}
		}
	}
	return mp
}

type Word struct {
	Key   string
	Count int
}
type WordsList []Word

func (w WordsList) Less(i, j int) bool { return w[i].Count < w[j].Count }
func (w WordsList) Swap(i, j int)      { w[i], w[j] = w[j], w[i] }
func (p WordsList) Len() int {
	return len(p)
}
func Top5Words(wordmap map[string]int) []Word {
	// Implement Me :)
	wordSlice := make(WordsList, len(wordmap))
	c := 0
	for k, w := range wordmap {
		wordSlice[c] = Word{k, w}
		c++
	}
	//fmt.Println(wordSlice)

	sort.Sort(sort.Reverse(wordSlice))
	return wordSlice[:5]
}

func main() {
	fmt.Println("Word Frequency Test")

	text := `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.`

	results := CountWords(text)
	//fmt.Println(results)
	MostCommon := Top5Words(results)

	fmt.Println(MostCommon)
}
