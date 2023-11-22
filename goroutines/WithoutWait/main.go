package main

import "fmt"

func main(){
	go some(1)
	go some(100)
	go some(400)
	go some(9)
	fmt.Println("Starting service!")
}
func some(v int){
	fmt.Println("Reached some:",v)
}