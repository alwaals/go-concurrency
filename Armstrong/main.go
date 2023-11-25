package main

import "fmt"

type MyInt int

func (npt *MyInt) IsArmstrong() bool {
	//res := 0
	n := *npt

	//n := *nptr

	n1 := float64(n/100)
	n2 := float64((n%100) / 10)
	n3 := float64(n % 10)

	fmt.Println(n1,n2,n3)
	// for int(val) > 0 {
	// 	rem := val % 10
	// 	val =  val/10
	// 	res = res + (int(rem) * int(rem) * int(rem))
	// }
	return false
}

func main() {
	fmt.Println("Armstrong Numbers")

	var num1 MyInt = 371
	fmt.Println(num1.IsArmstrong())
}
