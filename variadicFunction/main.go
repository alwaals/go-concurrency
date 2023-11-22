package main

import (
	"bytes"
	"fmt"
	"strings"
)

//Compare two slices of bytes
func main(){
	fmt.Println(someString("A","B","C"))
	fmt.Println(someString("ABC","NKJHB","oihihiC","Hello"))

	b1:=[]byte{'C','B','A'}
	b2:=[]byte{'C','B'}

	if res:=bytes.Compare(b1,b2);res==0{
		fmt.Println("Both are equal")
	}else{
		fmt.Println("Both are not equal")
	}
	var x int
	arr := []int{3, 5, 2}
	x, arr = arr[0], arr[1:]
	fmt.Println(x, arr)

}
func someString(str... string) string{
    return strings.Join(str,"")
}
