package main

import "fmt"

func main() {

	m := []int{2, 3, 4, 5}
	//ch := make(chan int)

	ch := readChannel(m)
	numCh := squareChannel(ch)

	for p := range numCh {
		fmt.Println(p)
	}

}
func readChannel(m []int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _,i := range m {
			//fmt.Println("readChannel:",i)
			ch <- i
		}
	}()
	return ch
}
func squareChannel(ch <-chan int) <-chan int {
	ch2 := make(chan int)
	go func() {
		defer close(ch2)
		for i := range ch {
			//fmt.Println("squareChannel:",i)
			ch2 <- i * i
		}
	}()
	return ch2
}
