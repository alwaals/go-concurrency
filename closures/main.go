package main

import "fmt"

func main(){

	closure1:=calculateClosure()
	closure2:=calculateClosure()
	fmt.Println(closure1(10))
	fmt.Println(closure2(20))


}
func calculateClosure() func(int) int{
	amount := 100

	debitAmount:=func(i int) int{
		amount-=i
		return amount
	}
	return debitAmount
}