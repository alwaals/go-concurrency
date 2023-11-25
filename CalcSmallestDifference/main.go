package main

import (
	"fmt"
	"sort"
)

func main() {

	arr1 := []int{9, 8, 7, 6}
	arr2 := []int{7, 3, 2, 5}

	fmt.Println(CalcSmallestDiff(arr1, arr2))
}
func CalcSmallestDiff(slice1, slice2 []int) int {

	sort.Ints(slice1)
	sort.Ints(slice2)
	fmt.Println(slice1, slice2)
	greater1 := slice1[len(slice1)-1]
	smaller2 := slice2[0]
	if greater1 >= smaller2 {
		return 0
	}
	return smaller2 - greater1

}
func CalcSmallestDifference(slice1, slice2 []int) int {
	min := 0
	sum := []int{}
	sum = append(sum, slice1...)
	sum = append(sum, slice2...)
	fmt.Println(sum)
	if len(sum) > 2 {
		min = diff(sum[1], sum[0])
		for i := 2; i < len(sum)-1; i++ {
			minSum := diff(sum[i], sum[i+1])
			if minSum < min {
				min = minSum
			}

		}
	} else if len(sum) == 2 {
		return sum[0] - sum[1]
	} else {
		return sum[0]
	}
	return min
}
func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
