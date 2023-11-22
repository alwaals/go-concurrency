package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	go func() {
		for i := 1; i <= 1000; i++ {
			wg.Add(1)
			go func(){
				defer wg.Done()
				getNum(ch)
			}()
		}
		wg.Wait()
		close(ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}
}
func getNum(ch chan int){
	time.Sleep(time.Second)
	ch<-rand.Intn(100)
}
