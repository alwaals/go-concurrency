package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)
	go doWork(ch)

	time.Sleep(time.Second * 3)
	close(ch)
}

func doWork(done <-chan string) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("doing work!")
		}
	}
}
