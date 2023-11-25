package main

import "fmt"

func main() {

	DoubleChar("gophers")
}
func DoubleChar(str string) string {
	// newStr := ""
	// for _, char := range str {
	// 	newStr = newStr + fmt.Sprintf("%c%c", char, char)
	// }
	runeStr := make([]rune, len(str)*2)
	for _, char := range str {
		runeStr = append(runeStr, char, char)
	}
	fmt.Println(string(runeStr))
	return string(runeStr)
}
