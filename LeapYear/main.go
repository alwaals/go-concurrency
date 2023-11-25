package main

import "fmt"

func main() {

	fmt.Println(LeapYear(1600))
}
func LeapYear(year int) bool {

	if year%4 == 0 {
		if !(year%100 == 0) {
			return true
		} else if year%400 == 0 {
			return true
		}
	}
	return false
}
