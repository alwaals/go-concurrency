package main

import (
	"fmt"
	"goroutines/stack/pop"
)

func main(){

	s:=[]string{"Hello","Swetha"}
	for len(s)>0{
		pop.Push(&s,"Nihira")
		fmt.Println("After Push:",s)
		s=pop.Pop(s)
		fmt.Println("After Pop:",s)
		break
	}
}