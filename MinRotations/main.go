package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println("Min Rotation Challenge")

	testArr := []int{15, 18, 2, 3, 6, 12}
	min := MinRotationsWithoutSort(testArr) // returns 2
	fmt.Println(min)

	testArr2 := []int{7, 9, 11, 12, 5}
	min2 := MinRotationsWithoutSort(testArr2) // return 4
	fmt.Println(min2)
}
func DiffSquares(m,n int)int{
	return int(math.Pow(float64(m),2)-math.Pow(float64(n),2))
}
func MinRotationsWithoutSort(inp []int) int {
	minIndex, maxIndex := 0, len(inp)-1
	counter := 0
	for i := 0; i < len(inp); i++ {
		//fmt.Println(inp[minIndex],inp[maxIndex])
		if inp[minIndex] > inp[maxIndex] {
			counter++
			minIndex++
			//maxIndex--
		} else {
			break
		}
	}

	return counter
}
func MinRotations(i []int) int {
	if len(i) > 0 {
		temp := make([]int, 0, len(i))
		temp = append(temp, i...)
		sort.Ints(temp)
		for i, v := range i {
			if temp[0] == v {
				return i
			}
		}
	}
	return 0
}
