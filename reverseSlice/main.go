package main

import "fmt"

//"welcome to strings"
func main(){
	sl:=[]int{3,2,1,9,8}
	fmt.Println(reverse(sl))
}
func reverse(sl []int) []int{
	newSl:=make([]int,len(sl))
	for a,b:=0,len(sl)-1;b>=len(sl)/2;a,b=a+1,b-1{
		newSl[a],newSl[b]=sl[b],sl[a]
	}
	return newSl
}
