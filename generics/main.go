package main

import "fmt"

func main(){

	fmt.Println(convertValue(10))
	fmt.Println(convertValue(10.9))

	mp:=make(map[string]int)
	mp["101"]=101

	k,v:=mp["102"]
	fmt.Println(k,v)

	var m interface{}

	m="str"

	p,s:=m.(string)

	fmt.Println(p,s)
}
type Numbers interface{
 int|float32|float64|int32|int64
}
func convertValue[T Numbers](t T) T{
    return t*t
}