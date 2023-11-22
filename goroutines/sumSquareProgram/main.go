package main

import "fmt"

// out : 2*2+3*3+4*4+5*5 =
var res []int

func main() {
	nums := []int{2, 3, 4, 5}

	ch := square(nums)
	//ch2 := sum(chS)

	for i := range ch {
		fmt.Println(i)
	}

}

func square(nums []int) <-chan int {
	chS := make(chan int)
	//chM := make(chan int)
	go func() {
		for _,p := range nums {
			//fmt.Println("num:",p)
			chS <- (p + p) +(p*p)

		}
		defer close(chS)

	}()
	return chS
}

// func sum(chS <-chan int) <-chan int {
// 	out := make(chan int)

// 	go func() {
// 		for i := range <-chS {
// 			out <- i
// 		}
// 		defer close(out)
// 	}()

// 	return out
// }
