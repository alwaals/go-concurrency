package main

import "fmt"

func main() {
	ch := make(chan string)
	ch2 := make(chan string)
	go func() {
		ch <- "data"
	}()
	go getData(ch2)

	for {
		select {
		case data := <-ch:
			fmt.Println(data)
		case data := <-ch2:
			fmt.Println(data)
		}
	}
}

func getData(ch chan string) {
	ch <- "data2"
}
