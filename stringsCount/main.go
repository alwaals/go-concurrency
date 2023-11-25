package main

import (
	"fmt"
	"sort"
	"strings"
)

type StringSlice struct {
	Key   string
	Count int
}
type MapList []StringSlice

// Welcome to the strings
func main() {

	str := "Welcome to the string prog"

	slices := strings.Fields(str)
	countStringsLen(slices)

}
func (m MapList) Less(i, j int) bool { return m[i].Count < m[j].Count }
func (m MapList) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m MapList) Len() int           { return len(m) }
func countStringsLen(slices []string){
    mp:= make(MapList,len(slices))
	c:=0
	for _,v:=range slices{
        mp[c]=StringSlice{v,len(v)}
		c++
	} 
	fmt.Println(mp)
	sort.Sort(sort.Reverse(mp))
	fmt.Println(mp)
}
func sortDescStrings(slices []string) {
	mp := make(map[string]int)
	for _, v := range slices {
		_, exists := mp[v]
		if !exists {
			mp[v] = 1
		} else {
			mp[v] += 1
		}
	}

	slice := make(MapList, len(mp))
	counter := 0
	for key, value := range mp {
		slice[counter] = StringSlice{key, value}
		counter++
	}
	fmt.Println("Before sorting:", slice)
	sort.Sort(sort.Reverse(slice))
	fmt.Println("After sorting:", slice)
}

