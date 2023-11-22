package main

import "fmt"

func main(){
	/// Shallow cloning
    i := []int{5,6}
	j := []int{1,2}
	temp:=i
	copy(i,j)
	fmt.Println(i,j,temp)
	
	/// Deep cloning
	p := []int{2,3}
	q := []int{4,5}
    t:=p
	p=q
	fmt.Println(p,q,t)
}