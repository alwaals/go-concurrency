package main

import "fmt"

func main() {
	s1 := "abc"
	s2 := "cba"
	isEqual := CheckPermutations(s1, s2)
	fmt.Println(isEqual)
}
func CheckPermutations(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	} else {
		mp := make(map[rune]int)
		for _, char := range s1 {
			_, exists := mp[char]
			if !exists {
				mp[char] = 1
			} else {
				mp[char] += 1
			}
		}
		for _, c := range s2 {
			v, exists := mp[c]
			if v == 1 {
				delete(mp, c)
			} else if exists {
				mp[c]--
			}
		}
		if len(mp) > 0 {
			return false
		}
	}
	return true
}
