package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

type MapList struct {
	Artist string
	Count  int
}
type MapListSlice []MapList

func (m MapListSlice) Less(i, j int) bool { return m[i].Count < m[j].Count }
func (m MapListSlice) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m MapListSlice) Len() int           { return len(m) }
func main() {

	f, err := os.Open("Top-50-musicality-global.csv")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	c := csv.NewReader(f)
	data := []string{}
	for {
		d, err := c.Read()
		if err != nil {
			if io.EOF == err {
				break
			}
			log.Fatal(err.Error())
		}
		data = append(data, d[3])
	}
	//fmt.Println(data,"\n",len(data))
	mp := make(map[string]int)
	//count:=0
	for _, v := range data {
		_, exist := mp[v]
		if exist {
			mp[v] += 1
		} else {
			mp[v] = 1
		}
	}
	// fmt.Println(mp, "\n", len(mp))

	//Sorting on slices

	freq := make(MapListSlice, len(mp))
	count := 0
	for i, v := range mp {
		freq[count] = MapList{i, v}
		count++
	}
	sort.Sort(sort.Reverse(freq))
	fmt.Println(freq)
}
